package jx3api

import "context"

type RankWantedResponse = RankResponse

// RankWanted 查询公开或私密的决斗挑战榜单。
func (c *Client) RankWanted(ctx context.Context, server string, mode int, token string) (*RankWantedResponse, error) {
	params := &struct {
		Server string `json:"server,omitempty"`
		Mode   int    `json:"mode"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Mode:   mode,
		Token:  token,
	}
	return requestData[RankWantedResponse](ctx, c, "/rank/wanted", params)
}
