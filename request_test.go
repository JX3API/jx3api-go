package jx3api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func testClient(t *testing.T, handler http.HandlerFunc) *Client {
	t.Helper()
	t.Setenv("JX3API_TOKEN", "")
	t.Setenv("JX3API_TICKET", "")
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)
	return NewClient(&Options{
		Token: "test-token", Ticket: "test-ticket",
		BaseURL: server.URL, HTTPClient: server.Client(),
	})
}

func TestRequestCredentials(t *testing.T) {
	for _, override := range []bool{false, true} {
		t.Run(map[bool]string{false: "defaults", true: "per-call"}[override], func(t *testing.T) {
			token, ticket := "test-token", "test-ticket"
			if override {
				token, ticket = "other-token", "other-ticket"
			}
			expectedToken, expectedTicket := token, ticket
			client := testClient(t, func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodPost || r.URL.Path != "/arena/recent" {
					t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
				}
				if r.Header.Get("Content-Type") != "application/json" || r.Header.Get("node") != "scf-node" {
					t.Error("missing request headers")
				}
				var body struct {
					Token  string `json:"token"`
					Ticket string `json:"ticket"`
				}
				if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
					t.Error(err)
				}
				if body.Token != expectedToken || body.Ticket != expectedTicket {
					t.Error("incorrect credential precedence")
				}
				w.Write([]byte(`{"status":200,"detail":{}}`))
			})
			client.Opts.Node = "scf-node"
			if !override {
				token, ticket = "", ""
			}
			_, err := client.ArenaRecent(context.Background(), "梦江南", "测试角色", 33, ticket, token)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestRequestErrors(t *testing.T) {
	cases := []struct {
		name   string
		http   int
		body   string
		status int
		text   string
	}{
		{"business", 200, `{"status":400,"detail":"暂无数据"}`, 400, "暂无数据"},
		{"upstream", 200, `{"status":502,"detail":"魔盒返回的角色 ID 不匹配"}`, 502, "角色 ID 不匹配"},
		{"http", 503, "Unavailable", 0, "HTTP 503"},
		{"invalid-json", 200, "<html>error</html>", 0, "invalid character"},
		{"unexpected-shape", 200, `[]`, 0, "cannot unmarshal"},
		{"missing-detail", 200, `{"status":200}`, 0, "status and detail are required"},
		{"missing-status", 200, `{"detail":{}}`, 0, "status and detail are required"},
		{"invalid-error-detail", 200, `{"status":400,"detail":{}}`, 0, "invalid API error detail"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			client := testClient(t, func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.http)
				w.Write([]byte(tc.body))
			})
			_, err := client.ActiveCalendar(context.Background(), 0)
			if err == nil || !strings.Contains(err.Error(), tc.text) {
				t.Fatalf("unexpected error: %v", err)
			}
			if tc.status != 0 {
				var apiError *APIError
				if !errors.As(err, &apiError) || apiError.Status != tc.status {
					t.Fatalf("API status was lost: %v", err)
				}
			}
		})
	}
}

func TestRawDetails(t *testing.T) {
	for _, detail := range []string{`[]`, `{}`, `""`, `0`, `false`, `null`} {
		t.Run(detail, func(t *testing.T) {
			client := testClient(t, func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(`{"status":200,"detail":` + detail + `}`))
			})
			data, err := requestData[json.RawMessage](context.Background(), client, "/test", nil)
			if err != nil || data == nil || string(*data) != detail {
				t.Fatalf("detail was not preserved: %s, %v", data, err)
			}
		})
	}
}

func TestEmptyList(t *testing.T) {
	client := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":200,"detail":{}}`))
	})
	data, err := client.SkillRework(context.Background())
	if err != nil || data == nil || len(*data) != 0 {
		t.Fatalf("unexpected empty list: %v, %v", data, err)
	}
}

func TestRequestCancellation(t *testing.T) {
	client := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		t.Error("a cancelled request reached the server")
	})
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := client.ActiveCalendar(ctx, 0)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("request cancellation was lost: %v", err)
	}
}
