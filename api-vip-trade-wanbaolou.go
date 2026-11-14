package jx3api

import (
	"context"
	"encoding/json"
)

type TradeWanbaolouResponse = map[string]json.RawMessage

// TradeWanbaolou 按万宝楼角色编号查询寄售角色详情。
func (c *Client) TradeWanbaolou(ctx context.Context, id string, token string) (*TradeWanbaolouResponse, error) {
	params := &struct {
		Id    string `json:"id"`
		Token string `json:"token,omitempty"`
	}{
		Id:    id,
		Token: token,
	}
	return requestData[TradeWanbaolouResponse](ctx, c, "/trade/wanbaolou", params)
}
