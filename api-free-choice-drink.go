package jx3api

import "context"

type ChoiceDrinkResponse []string

// ChoiceDrink 随机推荐 2～4 种饮品。
func (c *Client) ChoiceDrink(ctx context.Context) (*ChoiceDrinkResponse, error) {
	return requestData[ChoiceDrinkResponse](ctx, c, "/choice/drink", nil)
}
