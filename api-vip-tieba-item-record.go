package jx3api

import "context"

type TiebaItemRecordsResponse struct {
	Context string `json:"context"`
	Floor   int    `json:"floor"`
	Name    string `json:"name"`
	Reply   int64  `json:"reply"`
	Server  string `json:"server"`
	Time    int64  `json:"time"`
	Token   string `json:"token"`
	Url     int64  `json:"url"`
	Zone    string `json:"zone"`
}

// TiebaItemRecords 查询贴吧中的外观交易记录。
func (c *Client) TiebaItemRecords(ctx context.Context, server string, name string, limit int, token string) (*[]TiebaItemRecordsResponse, error) {
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
	return requestData[[]TiebaItemRecordsResponse](ctx, c, "/tieba/item/records", params)
}
