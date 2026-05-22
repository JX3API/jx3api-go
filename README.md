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

	daily, err := client.ActivateCalendar(context.TODO(), "梦江南", 0)
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
			"武林通鉴·公共任务：" + daily.Team[0],
			"武林通鉴·秘境任务：" + daily.Team[1],
			"武林通鉴·团队秘境：" + daily.Team[2],
		},
		"\n",
	)

	slog.Info(info)
}
```
