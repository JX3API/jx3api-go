package jx3api

import "context"

type RankLadderResponse = RankResponse

// RankLadder 查询跨服名剑榜单。
func (c *Client) RankLadder(ctx context.Context, server string, mode int, token string) (*RankLadderResponse, error) {
	params := &struct {
		Server string `json:"server,omitempty"`
		Mode   int    `json:"mode"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Mode:   mode,
		Token:  token,
	}
	return requestData[RankLadderResponse](ctx, c, "/rank/ladder", params)
}
