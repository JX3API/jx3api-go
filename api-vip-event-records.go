package jx3api

import "context"

type EventRecordsResponse struct {
	Zone   string `json:"zone"`
	Server string `json:"server"`
	Name   string `json:"name"`
	Event  string `json:"event"`
	Level  int    `json:"level"`
	Status *int   `json:"status"`
	Source *int   `json:"source"`
	Time   *int64 `json:"time"`
}

// EventRecords 查询角色的奇遇触发记录。
func (c *Client) EventRecords(ctx context.Context, server string, name string, full int, token string) (*[]EventRecordsResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Name   string `json:"name"`
		Full   int    `json:"full"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Name:   name,
		Full:   full,
		Token:  token,
	}
	return requestData[[]EventRecordsResponse](ctx, c, "/event/records", params)
}
