package jx3api

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestCalendarModes(t *testing.T) {
	client := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		var params map[string]any
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Error(err)
		}
		if r.URL.Path != "/active/calendar" || params["server"] != nil {
			t.Error("obsolete calendar path or server parameter")
		}
		if params["mode"] == "list" {
			w.Write([]byte(`{"status":200,"detail":{"today":{"date":"2026-11-14"},"total":[{"war":"大战"}]}}`))
		} else {
			w.Write([]byte(`{"status":200,"detail":{"date":"2026-11-14","draw":["美人图"],"lucky":["奇遇"],"weekly":{"conn":["公共周常"],"raid":["团队周常"]}}}`))
		}
	})
	daily, err := client.ActivateCalendar(context.Background(), "梦江南", 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(daily.Draw) != 1 || len(daily.Lucky) != 1 || len(daily.Weekly.Conn) != 1 {
		t.Fatalf("calendar fields were lost: %+v", daily)
	}
	list, err := client.ActivateListCalendar(context.Background(), 1)
	if err != nil || list == nil || len(list.Total) != 1 || list.Total[0].War != "大战" {
		t.Fatalf("calendar list was not decoded: %+v, %v", list, err)
	}
}

func TestSchoolResponses(t *testing.T) {
	responses := map[string]string{
		"/school/matrix": `[{"id":230,"name":"七绝逍遥阵","levels":[{"level":1,"desc":"阵法效果"}]}]`,
		"/school/skills": `[{"id":1,"name":"阳明指","interval":1.5,"distance":20,"passive":false}]`,
		"/school/talent": `[{"slot":1,"pointId":61,"options":[{"id":14635,"name":"踏莲","passive":true}]}]`,
	}
	client := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		var params map[string]any
		json.NewDecoder(r.Body).Decode(&params)
		if params["kungfu"] != "花间游" || params["name"] != nil || params["ticket"] != nil || r.Header.Get("ticket") != "" {
			t.Error("incorrect school parameters or obsolete ticket requirement")
		}
		w.Write([]byte(`{"status":200,"detail":` + responses[r.URL.Path] + `}`))
	})
	client.Opts.Ticket = ""
	matrix, err := client.SchoolMatrix(context.Background(), "花间游", "")
	if err != nil || matrix == nil || len(*matrix) != 1 || len((*matrix)[0].Levels) != 1 {
		t.Fatalf("incorrect matrix response: %v, %v", matrix, err)
	}
	skills, err := client.SchoolSkills(context.Background(), "花间游", "")
	if err != nil || skills == nil || len(*skills) != 1 || (*skills)[0].Interval != 1.5 {
		t.Fatalf("incorrect skill response: %v, %v", skills, err)
	}
	talents, err := client.SchoolTalent(context.Background(), "花间游", "")
	if err != nil || talents == nil || len(*talents) != 1 || len((*talents)[0].Options) != 1 {
		t.Fatalf("incorrect talent response: %v, %v", talents, err)
	}
}

func TestRankResponseVariants(t *testing.T) {
	for _, detail := range []string{
		`{"server":"梦江南","label":"个人排行","data":[{"dwID":9007199254740993}]}`,
		`[{"server":"梦江南","label":"个人排行","data":[{"dwID":9007199254740993}]}]`,
	} {
		t.Run(detail[:1], func(t *testing.T) {
			client := testClient(t, func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/rank/records" {
					t.Error("obsolete rank path")
				}
				w.Write([]byte(`{"status":200,"detail":` + detail + `}`))
			})
			data, err := client.RankStatistical(context.Background(), "梦江南", "老江湖五十强", "")
			if err != nil || data == nil || len(*data) != 1 {
				t.Fatalf("rank response was not normalized: %v, %v", data, err)
			}
			if string((*data)[0].Data[0]["dwID"]) != "9007199254740993" {
				t.Fatal("rank record lost integer precision")
			}
		})
	}
}

func TestStatusResponseVariants(t *testing.T) {
	for _, detail := range []string{
		`{"zone":"电信区","server":"梦江南","status":"正常"}`,
		`[{"zone":"电信区","server":"梦江南","status":1,"lasttime":1794614400,"shuttime":0}]`,
	} {
		client := testClient(t, func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/server/status/check" {
				t.Error("obsolete server status path")
			}
			w.Write([]byte(`{"status":200,"detail":` + detail + `}`))
		})
		data, err := client.StatusCheck(context.Background(), 1, "梦江南")
		if err != nil || data == nil || len(*data) != 1 || (*data)[0].Server != "梦江南" {
			t.Fatalf("incorrect status response: %v, %v", data, err)
		}
	}
}

func TestArenaEmptyPerformances(t *testing.T) {
	client := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":200,"detail":{"campId":1,"performance":{"2v2":[],"3v3":[],"5v5":{"mmr":2200}}}}`))
	})
	data, err := client.ArenaRecent(context.Background(), "梦江南", "测试角色", 0, "", "")
	if err != nil || data == nil || data.CampId != 1 || data.Performance.V5v5.Mmr != 2200 {
		t.Fatalf("incorrect arena response: %v, %v", data, err)
	}
}

func TestNullableEventFields(t *testing.T) {
	client := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":200,"detail":[{"name":"测试角色","event":"三山四海","level":1,"status":null,"source":null,"time":null}]}`))
	})
	data, err := client.EventRecords(context.Background(), "梦江南", "测试角色", 1, "")
	if err != nil || data == nil || len(*data) != 1 {
		t.Fatalf("incorrect event response: %v, %v", data, err)
	}
	if event := (*data)[0]; event.Time != nil || event.Status != nil || event.Source != nil {
		t.Fatal("null event values were not preserved")
	}
}
