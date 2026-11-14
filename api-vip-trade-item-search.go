package jx3api

import "context"

type TradeItemSearchResponse struct {
	Alias    string  `json:"alias"`
	Category string  `json:"category"`
	Date     *string `json:"date"`
	Desc     *string `json:"desc"`
	Name     string  `json:"name"`
	Retail   int     `json:"retail"`
	View     *string `json:"view"`
}

// TradeItemSearch 搜索外观及其基础资料。
func (c *Client) TradeItemSearch(ctx context.Context, name string, token string) (*[]TradeItemSearchResponse, error) {
	params := &struct {
		Name  string `json:"name"`
		Token string `json:"token,omitempty"`
	}{
		Name:  name,
		Token: token,
	}
	return requestData[[]TradeItemSearchResponse](ctx, c, "/trade/item/search", params)
}
