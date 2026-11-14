package jx3api

import "context"

type SaohuaContentResponse struct {
	Text string `json:"text"`
}

// SaohuaContent 随机获取一条舔狗日志。
func (c *Client) SaohuaContent(ctx context.Context) (*SaohuaContentResponse, error) {
	return requestData[SaohuaContentResponse](ctx, c, "/saohua/content", nil)
}
