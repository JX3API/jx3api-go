package jx3api

import "context"

type SaohuaZhananResponse struct {
	Text string `json:"text"`
}

// SaohuaZhanan 随机获取一条渣男语录。
func (c *Client) SaohuaZhanan(ctx context.Context) (*SaohuaZhananResponse, error) {
	return requestData[SaohuaZhananResponse](ctx, c, "/saohua/zhanan", nil)
}
