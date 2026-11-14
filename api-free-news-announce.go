package jx3api

import (
	"context"
	"encoding/json"
)

type NewsAnnounceResponse struct {
	Catid string `json:"catid"`
	Date  string `json:"date"`
	Desc  struct {
		Asktime   string          `json:"asktime"`
		Content   string          `json:"content"`
		ExtData   json.RawMessage `json:"ext_data"`
		Inputtime int64           `json:"inputtime"`
		Title     string          `json:"title"`
		Top       string          `json:"top"`
	} `json:"desc"`
	Title string `json:"title"`
	Type  string `json:"type"`
	Url   string `json:"url"`
}

// NewsAnnounce 查询游戏维护更新公告。
func (c *Client) NewsAnnounce(ctx context.Context, limit int) (*[]NewsAnnounceResponse, error) {
	params := &struct {
		Limit int `json:"limit"`
	}{
		Limit: limit,
	}
	return requestData[[]NewsAnnounceResponse](ctx, c, "/news/announce", params)
}
