package jx3api

import "context"

type CardRecordResponse = CardCachedResponse

// CardRecord 查询角色名片。
func (c *Client) CardRecord(ctx context.Context, server string, name string, token string) (*CardRecordResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Name   string `json:"name"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Name:   name,
		Token:  token,
	}
	return requestData[CardRecordResponse](ctx, c, "/card/record", params)
}
