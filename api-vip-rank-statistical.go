package jx3api

import "context"

type RankRecordsResponse = RankResponse

type RankStatisticalResponse = RankRecordsResponse

// RankRecords 查询个人排行与贡献榜。
func (c *Client) RankRecords(ctx context.Context, server string, name string, token string) (*RankRecordsResponse, error) {
	params := &struct {
		Server string `json:"server,omitempty"`
		Name   string `json:"name"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Name:   name,
		Token:  token,
	}
	return requestData[RankRecordsResponse](ctx, c, "/rank/records", params)
}

// RankStatistical is an alias for RankRecords.
func (c *Client) RankStatistical(ctx context.Context, server string, name string, token string) (*RankRecordsResponse, error) {
	return c.RankRecords(ctx, server, name, token)
}
