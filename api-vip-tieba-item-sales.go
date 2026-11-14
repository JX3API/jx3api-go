package jx3api

import "context"

type TiebaItemSalesResponse struct {
	Server []struct {
		Date   string `json:"date"`
		Sale   int    `json:"sale"`
		Server string `json:"server"`
		Value  int    `json:"value"`
		Zone   string `json:"zone"`
	} `json:"server"`
	Tieba []struct {
		Date   string `json:"date"`
		Sale   int    `json:"sale"`
		Server string `json:"server"`
		Value  int    `json:"value"`
		Zone   string `json:"zone"`
	} `json:"tieba"`
}

// TiebaItemSales 查询物品在贴吧及其他渠道的最近价格记录。
func (c *Client) TiebaItemSales(ctx context.Context, name string, token string) (*TiebaItemSalesResponse, error) {
	params := &struct {
		Name  string `json:"name"`
		Token string `json:"token,omitempty"`
	}{
		Name:  name,
		Token: token,
	}
	return requestData[TiebaItemSalesResponse](ctx, c, "/tieba/item/sales", params)
}
