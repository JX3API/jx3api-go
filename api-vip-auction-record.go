package jx3api

import "context"

type AuctionRecordsResponse struct {
	CampName   string `json:"campName"`
	ItemAmount string `json:"itemAmount"`
	ItemName   string `json:"itemName"`
	RoleName   string `json:"roleName"`
	Server     string `json:"server"`
	Time       int64  `json:"time"`
	Zone       string `json:"zone"`
}

type AuctionRecordResponse = AuctionRecordsResponse

// AuctionRecords 查询阵营拍卖成交记录。
func (c *Client) AuctionRecords(ctx context.Context, server string, name string, limit int, token string) (*[]AuctionRecordsResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Name   string `json:"name,omitempty"`
		Limit  int    `json:"limit"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Name:   name,
		Limit:  limit,
		Token:  token,
	}
	return requestData[[]AuctionRecordsResponse](ctx, c, "/auction/records", params)
}

// AuctionRecord is an alias for AuctionRecords.
func (c *Client) AuctionRecord(ctx context.Context, server string, name string, limit int, token string) (*[]AuctionRecordsResponse, error) {
	return c.AuctionRecords(ctx, server, name, limit, token)
}
