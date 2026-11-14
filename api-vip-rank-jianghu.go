package jx3api

import "context"

type RankJianghuResponse = RankResponse

// RankJianghu 查询江湖浪客与捕快荣誉榜。
func (c *Client) RankJianghu(ctx context.Context, server string, name string, token string) (*RankJianghuResponse, error) {
	params := &struct {
		Server string `json:"server,omitempty"`
		Name   string `json:"name"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Name:   name,
		Token:  token,
	}
	return requestData[RankJianghuResponse](ctx, c, "/rank/jianghu", params)
}
