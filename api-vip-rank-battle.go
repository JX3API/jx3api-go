package jx3api

import "context"

type RankBattleResponse = RankResponse

// RankBattle 查询绝境战场榜单。
func (c *Client) RankBattle(ctx context.Context, server string, name string, token string) (*RankBattleResponse, error) {
	params := &struct {
		Server string `json:"server,omitempty"`
		Name   string `json:"name"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Name:   name,
		Token:  token,
	}
	return requestData[RankBattleResponse](ctx, c, "/rank/battle", params)
}
