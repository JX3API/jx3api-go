package jx3api

import "context"

type EventStrategyResponse struct {
	Box   int    `json:"box"`
	Level int    `json:"level"`
	Name  string `json:"name"`
	Type  string `json:"type"`
}

// EventStrategy 查询奇遇的分类和等级资料。
func (c *Client) EventStrategy(ctx context.Context, name string, token string) (*EventStrategyResponse, error) {
	params := &struct {
		Name  string `json:"name"`
		Token string `json:"token,omitempty"`
	}{
		Name:  name,
		Token: token,
	}
	return requestData[EventStrategyResponse](ctx, c, "/event/strategy", params)
}
