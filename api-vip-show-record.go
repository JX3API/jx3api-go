package jx3api

import "context"

type FireworkRecordsResponse struct {
	Firework string `json:"firework"`
	MapName  string `json:"mapName"`
	Matched  int    `json:"matched"`
	Receiver string `json:"receiver"`
	Sender   string `json:"sender"`
	Server   string `json:"server"`
	Status   int    `json:"status"`
	Time     int64  `json:"time"`
	Zone     string `json:"zone"`
}

type ShowRecordResponse = FireworkRecordsResponse

// FireworkRecords 查询角色发送或接收的烟花记录。
func (c *Client) FireworkRecords(ctx context.Context, server string, name string, token string) (*[]FireworkRecordsResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Name   string `json:"name"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Name:   name,
		Token:  token,
	}
	return requestData[[]FireworkRecordsResponse](ctx, c, "/firework/records", params)
}

// ShowRecord is an alias for FireworkRecords.
func (c *Client) ShowRecord(ctx context.Context, server string, name string, token string) (*[]FireworkRecordsResponse, error) {
	return c.FireworkRecords(ctx, server, name, token)
}
