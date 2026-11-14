package jx3api

import "context"

type TradeDemonResponse map[string]*string

// TradeDemon 查询各渠道的金币价格。
func (c *Client) TradeDemon(ctx context.Context, server string, limit int, token string) (*[]TradeDemonResponse, error) {
	params := &struct {
		Server string `json:"server,omitempty"`
		Limit  int    `json:"limit"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Limit:  limit,
		Token:  token,
	}
	return requestData[[]TradeDemonResponse](ctx, c, "/trade/demon", params)
}
