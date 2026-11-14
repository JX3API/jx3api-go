package jx3api

import "context"

type RankContestResponse = RankResponse

// RankContest 查询指定阵营的武林争霸榜单。
func (c *Client) RankContest(ctx context.Context, server string, camp int, token string) (*RankContestResponse, error) {
	params := &struct {
		Server string `json:"server,omitempty"`
		Camp   int    `json:"camp"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Camp:   camp,
		Token:  token,
	}
	return requestData[RankContestResponse](ctx, c, "/rank/contest", params)
}
