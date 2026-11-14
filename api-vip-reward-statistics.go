package jx3api

import "context"

type RewardStatisticsResponse struct {
	ItemAmount int    `json:"itemAmount"`
	ItemName   string `json:"itemName"`
	MapName    string `json:"mapName"`
	RoleName   string `json:"roleName"`
	Server     string `json:"server"`
	Time       int64  `json:"time"`
	Zone       string `json:"zone"`
}

// RewardStatistics 查询物品的掉落记录。
func (c *Client) RewardStatistics(ctx context.Context, server string, name string, limit int, token string) (*[]RewardStatisticsResponse, error) {
	params := &struct {
		Server string `json:"server,omitempty"`
		Name   string `json:"name"`
		Limit  int    `json:"limit"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Name:   name,
		Limit:  limit,
		Token:  token,
	}
	return requestData[[]RewardStatisticsResponse](ctx, c, "/reward/statistics", params)
}
