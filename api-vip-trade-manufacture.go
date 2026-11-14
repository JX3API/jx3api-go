package jx3api

import (
	"context"
	"encoding/json"
)

type TradeManufactureResponse = map[string]json.RawMessage

// TradeManufacture 查询成品售价并计算配方材料成本。
func (c *Client) TradeManufacture(ctx context.Context, server string, name string, source int, token string) (*TradeManufactureResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Name   string `json:"name"`
		Source int    `json:"source"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Name:   name,
		Source: source,
		Token:  token,
	}
	return requestData[TradeManufactureResponse](ctx, c, "/trade/manufacture", params)
}
