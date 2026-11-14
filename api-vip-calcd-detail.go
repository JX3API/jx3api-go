package jx3api

import "context"

type CalcdDetailResponse = CalcdAttributeResponse

// CalcdDetail 查询角色指定心法及类型的已收录配装。
func (c *Client) CalcdDetail(ctx context.Context, server string, name string, kungfu string, typ string, format string, token string) (*CalcdDetailResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Name   string `json:"name"`
		Kungfu string `json:"kungfu"`
		Type   string `json:"type"`
		Format string `json:"format,omitempty"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Name:   name,
		Kungfu: kungfu,
		Type:   typ,
		Format: format,
		Token:  token,
	}
	return requestData[CalcdDetailResponse](ctx, c, "/calcd/detail", params)
}
