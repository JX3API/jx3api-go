package jx3api

import "context"

type WickedRecordsResponse struct {
	MapName string `json:"mapName"`
	Server  string `json:"server"`
	Time    int64  `json:"time"`
	Zone    string `json:"zone"`
}

type SmiteRecord = WickedRecordsResponse

// WickedRecords 查询诛恶事件刷新记录。
func (c *Client) WickedRecords(ctx context.Context, server string, limit int, token string) (*[]WickedRecordsResponse, error) {
	params := &struct {
		Server string `json:"server,omitempty"`
		Limit  int    `json:"limit"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Limit:  limit,
		Token:  token,
	}
	return requestData[[]WickedRecordsResponse](ctx, c, "/wicked/records", params)
}

// SmiteRecords is an alias for WickedRecords.
func (c *Client) SmiteRecords(ctx context.Context, token string) (*[]WickedRecordsResponse, error) {
	return c.WickedRecords(ctx, "", 20, token)
}
