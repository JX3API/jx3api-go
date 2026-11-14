package jx3api

import "context"

type ShareStayResponse struct {
	Data struct {
		MapName string  `json:"map_name"`
		MaxTime float64 `json:"max_time"`
		MinTime int64   `json:"min_time"`
		NpcName string  `json:"npc_name"`
	} `json:"data"`
	Name   string `json:"name"`
	Server string `json:"server"`
	Time   int64  `json:"time"`
	Zone   string `json:"zone"`
}

// ShareStay 查询宠物最近触发时间及刷新信息。
func (c *Client) ShareStay(ctx context.Context, server string, token string) (*[]ShareStayResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Token:  token,
	}
	return requestData[[]ShareStayResponse](ctx, c, "/share/stay", params)
}
