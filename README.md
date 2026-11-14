# jx3api-go
The Golang SDK to the [JX3API](https://www.jx3api.com).

## Installation
```bash
go get -u github.com/JX3API/jx3api-go@latest
```

## Quick Start
```golang
package main

import (
	"context"
	"log/slog"
	"strings"

	"github.com/JX3API/jx3api-go"
)

func main() {
	opts := &jx3api.Options{
		Token:  "put your token here",
		Ticket: "put your ticket here",
	}

	client := jx3api.NewClient(opts)

	daily, err := client.ActiveCalendar(context.Background(), 0)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	info := strings.Join(
		[]string{
			"当前时间：" + daily.Date,
			"秘境大战：" + daily.War,
			"战场任务：" + daily.Battle,
			"宗门事件：" + daily.School,
			"阵营任务：" + daily.Orecar,
			"武林通鉴·公共任务：" + strings.Join(daily.Weekly.Conn, "、"),
			"武林通鉴·团队秘境：" + strings.Join(daily.Weekly.Raid, "、"),
		},
		"\n",
	)

	slog.Info(info)
}
```

## API

Endpoints follow the [current OpenAPI specification](https://www.jx3api.com/openapi).
All requests use the same host and POST JSON. The shared requester checks `status`
and decodes `detail` directly into each method's response type.

`Options.Token` and `Options.Ticket` fall back to `JX3API_TOKEN` and `JX3API_TICKET`.
Pass an empty credential argument to use those defaults; a nonempty per-call
credential takes precedence. `Options.Node` supplies the optional SCF node header.
`Options.HTTPClient` and `Options.BaseURL` support timeouts and custom transports.
Credentials are never logged by the requester.

```go
skills, err := client.SchoolSkills(ctx, "花间游", "")
role, err := client.RoleDetail(ctx, server, name, 1, "")
equipment, err := client.CalcdAttribute(ctx, server, name, "calcd", "")
achievement, err := client.RoleAchievement(ctx, server, name, 3, "", "")
```

- `SchoolMatrix`, `SchoolSkills`, `SchoolTalent`, and `RankTrials` use `kungfu`.
  The three school methods no longer take a ticket and return slices.
- `MentorSearch` and `RecruitSearch` send `label` and accept `limit`; their results
  are slices rather than objects containing `data`.
- `RoleDetail` accepts `history`, `EventRecords` accepts `full`, `EventRecent`
  accepts `limit`, `CardRecords` accepts `like`, and `RanchRecords` accepts `expired`.
- Daily calendars expose `Lucky`, `Draw`, and `Weekly`; calendar lists use `Total`.
- Rankings and server statuses normalize single-server objects and all-server
  arrays into slices. Ranking columns retain their original JSON names and values.
- Nullable event and horse fields use pointers. Equipment models accept both
  `calcd` and raw formats; raw and extensible fields retain `json.RawMessage`.
- `APIError` exposes the API's `Status` and `Detail` and supports `errors.As`.
  HTTP, transport, and JSON decoding errors are also returned to callers.

Compatible aliases remain for `ActivateCalendar`, `ActivateListCalendar`,
`ExamAnswer`, `SchoolFoods`, `ServerMaster`, `StatusCheck`, `NewsAllnews`,
`ActiveMonster`, `DuowanStatistics`, `MechCalculator`, `RankStatistical`,
`RoleMonster`, `ShowRecord`, and `SmiteRecords`.

Removed endpoints: `HomeFurniture`, `HomeTravel`, `NextEvent`, `EventUnfinished`,
`FenxianRecords`, and `MineCart`.

Fields without a fixed response schema are exposed as `json.RawMessage` or maps.

## Tests

```bash
go test ./...
```

The regression tests use local HTTP fixtures and do not require API credentials.
