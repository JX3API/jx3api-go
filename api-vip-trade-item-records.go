package jx3api

import "context"

type TradeItemRecordsResponse struct {
	Alias    string `json:"alias"`
	Category string `json:"category"`
	Date     string `json:"date"`
	Desc     string `json:"desc"`
	List     []struct {
		List []struct {
			Date          string  `json:"date"`
			Id            *string `json:"id,omitempty"`
			Index         *string `json:"index,omitempty"`
			RemainingTime *int64  `json:"remainingTime,omitempty"`
			Sale          int     `json:"sale"`
			Server        string  `json:"server"`
			Source        *int    `json:"source,omitempty"`
			Status        *int    `json:"status,omitempty"`
			Token         *string `json:"token,omitempty"`
			Value         int     `json:"value"`
			Zone          string  `json:"zone"`
		} `json:"list"`
		Name string `json:"name"`
	} `json:"list"`
	Name   string `json:"name"`
	Retail int    `json:"retail"`
	View   string `json:"view"`
}

// TradeItemRecords 按物品全名查询成交价和寄售挂单。
func (c *Client) TradeItemRecords(ctx context.Context, server string, name string, token string) (*TradeItemRecordsResponse, error) {
	params := &struct {
		Server string `json:"server,omitempty"`
		Name   string `json:"name"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Name:   name,
		Token:  token,
	}
	return requestData[TradeItemRecordsResponse](ctx, c, "/trade/item/records", params)
}
