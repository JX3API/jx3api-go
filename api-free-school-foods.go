package jx3api

import "context"

type FoodListResponse struct {
	Boost  string `json:"boost"`
	Class  string `json:"class"`
	Color  string `json:"color"`
	Kungfu string `json:"kungfu"`
	Name   string `json:"name"`
	School string `json:"school"`
}

type SchoolFoodsResponse = FoodListResponse

// FoodList 查询心法的小吃小药。
func (c *Client) FoodList(ctx context.Context, name string) (*[]FoodListResponse, error) {
	params := &struct {
		Name string `json:"name,omitempty"`
	}{
		Name: name,
	}
	return requestData[[]FoodListResponse](ctx, c, "/food/list", params)
}

// SchoolFoods is an alias for FoodList.
func (c *Client) SchoolFoods(ctx context.Context) (*[]FoodListResponse, error) {
	return c.FoodList(ctx, "")
}
