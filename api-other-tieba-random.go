package jx3api

import "context"

type TiebaRandomResponse struct {
	Date   string `json:"date"`
	Name   string `json:"name"`
	Server string `json:"server"`
	Tags   string `json:"tags"`
	Title  string `json:"title"`
	Url    int64  `json:"url"`
	Zone   string `json:"zone"`
}

// TiebaRandom 按类型随机查询近 31 天的贴吧帖子。
func (c *Client) TiebaRandom(ctx context.Context, tags string, server string, limit int, token string) (*[]TiebaRandomResponse, error) {
	params := &struct {
		Tags   string `json:"tags"`
		Server string `json:"server,omitempty"`
		Limit  int    `json:"limit"`
		Token  string `json:"token,omitempty"`
	}{
		Tags:   tags,
		Server: server,
		Limit:  limit,
		Token:  token,
	}
	return requestData[[]TiebaRandomResponse](ctx, c, "/tieba/random", params)
}
