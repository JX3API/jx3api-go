package jx3api

import "context"

type CardDetailResponse struct {
	PraiseTotalCount  int    `json:"praiseTotalCount"`
	ShowCardPresetUrl string `json:"showCardPresetUrl"`
}

// CardDetail 查询角色名片图片与获赞数。
func (c *Client) CardDetail(ctx context.Context, server string, name string, ticket string, token string) (*CardDetailResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Name   string `json:"name"`
		Ticket string `json:"ticket,omitempty"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Name:   name,
		Ticket: ticket,
		Token:  token,
	}
	return requestData[CardDetailResponse](ctx, c, "/card/detail", params)
}
