package jx3api

import "context"

type SaohuaRandomResponse struct {
	Text string `json:"text"`
}

// SaohuaRandom 随机获取一条骚话。
func (c *Client) SaohuaRandom(ctx context.Context) (*SaohuaRandomResponse, error) {
	return requestData[SaohuaRandomResponse](ctx, c, "/saohua/random", nil)
}
