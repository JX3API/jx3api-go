package jx3api

import "context"

type EventStatisticsResponse struct {
	Event  string `json:"event"`
	Name   string `json:"name"`
	Server string `json:"server"`
	Source int    `json:"source"`
	Status int    `json:"status"`
	Time   int64  `json:"time"`
	Zone   string `json:"zone"`
}

// EventStatistics 查询指定奇遇最近的触发记录。
func (c *Client) EventStatistics(ctx context.Context, server string, name string, limit int, token string) (*[]EventStatisticsResponse, error) {
	params := &struct {
		Server string `json:"server,omitempty"`
		Name   string `json:"name"`
		Limit  int    `json:"limit"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Name:   name,
		Limit:  limit,
		Token:  token,
	}
	return requestData[[]EventStatisticsResponse](ctx, c, "/event/statistics", params)
}
