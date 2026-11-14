package jx3api

import "context"

type RankGuildResponse = RankResponse

// RankGuild 查询帮会排行与战功榜。
func (c *Client) RankGuild(ctx context.Context, server string, name string, token string) (*RankGuildResponse, error) {
	params := &struct {
		Server string `json:"server,omitempty"`
		Name   string `json:"name"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Name:   name,
		Token:  token,
	}
	return requestData[RankGuildResponse](ctx, c, "/rank/guild", params)
}
