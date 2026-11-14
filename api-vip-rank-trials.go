package jx3api

import "context"

type RankTrialsResponse = RankResponse

// RankTrials 查询指定心法的试炼之地榜单。
func (c *Client) RankTrials(ctx context.Context, server string, kungfu string, token string) (*RankTrialsResponse, error) {
	params := &struct {
		Server string `json:"server,omitempty"`
		Kungfu string `json:"kungfu"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Kungfu: kungfu,
		Token:  token,
	}
	return requestData[RankTrialsResponse](ctx, c, "/rank/trials", params)
}
