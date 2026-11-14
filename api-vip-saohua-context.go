package jx3api

import "context"

type SaohuaContextResponse struct {
	Content string `json:"content"`
}

// SaohuaContext 按分类随机获取一条语录。
func (c *Client) SaohuaContext(ctx context.Context, name string, token string) (*SaohuaContextResponse, error) {
	params := &struct {
		Name  string `json:"name"`
		Token string `json:"token,omitempty"`
	}{
		Name:  name,
		Token: token,
	}
	return requestData[SaohuaContextResponse](ctx, c, "/saohua/context", params)
}
