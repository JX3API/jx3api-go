package jx3api

import "context"

type EventCollectResponse struct {
	Count int    `json:"count"`
	Event string `json:"event"`
	Last  struct {
		Name string `json:"name"`
		Time int64  `json:"time"`
	} `json:"last"`
	Server string `json:"server"`
	Zone   string `json:"zone"`
}

// EventCollect 汇总服务器近期的奇遇次数及最近触发记录。
func (c *Client) EventCollect(ctx context.Context, server string, num int, token string) (*[]EventCollectResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Num    int    `json:"num"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Num:    num,
		Token:  token,
	}
	return requestData[[]EventCollectResponse](ctx, c, "/event/collect", params)
}
