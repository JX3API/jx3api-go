package jx3api

import "context"

type EventRecentResponse struct {
	Event  string `json:"event"`
	Name   string `json:"name"`
	Server string `json:"server"`
	Source int    `json:"source"`
	Status int    `json:"status"`
	Time   int64  `json:"time"`
	Zone   string `json:"zone"`
}

// EventRecent 查询服务器最近的奇遇触发记录。
func (c *Client) EventRecent(ctx context.Context, server string, limit int, token string) (*[]EventRecentResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Limit  int    `json:"limit"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Limit:  limit,
		Token:  token,
	}
	return requestData[[]EventRecentResponse](ctx, c, "/event/recent", params)
}
