package jx3api

import "context"

type CardRecordsResponse = CardCachedResponse

// CardRecords 查询角色的历史名片。
func (c *Client) CardRecords(ctx context.Context, server string, name string, like bool, token string) (*[]CardRecordsResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Name   string `json:"name"`
		Like   bool   `json:"like"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Name:   name,
		Like:   like,
		Token:  token,
	}
	return requestData[[]CardRecordsResponse](ctx, c, "/card/records", params)
}
