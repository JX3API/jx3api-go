package jx3api

import "context"

type RankCastleResponse = RankResponse

// RankCastle 查询据点贡献榜。
func (c *Client) RankCastle(ctx context.Context, server string, name string, token string) (*RankCastleResponse, error) {
	params := &struct {
		Server string `json:"server,omitempty"`
		Name   string `json:"name"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Name:   name,
		Token:  token,
	}
	return requestData[RankCastleResponse](ctx, c, "/rank/castle", params)
}
