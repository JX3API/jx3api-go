package jx3api

import "context"

type ChoiceFoodResponse []string

// ChoiceFood 随机推荐 2～4 种食物。
func (c *Client) ChoiceFood(ctx context.Context) (*ChoiceFoodResponse, error) {
	return requestData[ChoiceFoodResponse](ctx, c, "/choice/food", nil)
}
