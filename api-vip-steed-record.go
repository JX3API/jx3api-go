package jx3api

import "context"

type SteedRecordsResponse struct {
	AuctionAmount   *string `json:"auctionAmount,omitempty"`
	AuctionCampName *string `json:"auctionCampName,omitempty"`
	AuctionRoleName *string `json:"auctionRoleName,omitempty"`
	AuctionTime     *int64  `json:"auctionTime,omitempty"`
	CaptureCampName *string `json:"captureCampName,omitempty"`
	CaptureRoleName *string `json:"captureRoleName,omitempty"`
	CaptureTime     *int64  `json:"captureTime,omitempty"`
	Data            *struct {
		AuctionAmount   *string `json:"auctionAmount,omitempty"`
		AuctionCampName *string `json:"auctionCampName,omitempty"`
		AuctionRoleName *string `json:"auctionRoleName,omitempty"`
		AuctionTime     *int64  `json:"auctionTime,omitempty"`
		CaptureCampName string  `json:"captureCampName"`
		CaptureRoleName string  `json:"captureRoleName"`
		CaptureTime     int64   `json:"captureTime"`
		MapName         string  `json:"mapName"`
		RefreshTime     int64   `json:"refreshTime"`
		Server          string  `json:"server"`
		Zone            string  `json:"zone"`
	} `json:"data,omitempty"`
	MapName     *string `json:"mapName,omitempty"`
	RefreshTime *int64  `json:"refreshTime,omitempty"`
	Server      string  `json:"server"`
	Zone        *string `json:"zone,omitempty"`
}

type SteedRecordResponse = SteedRecordsResponse

// SteedRecords 查询的卢出世、捕获与拍卖记录。
func (c *Client) SteedRecords(ctx context.Context, server string, token string) (*[]SteedRecordsResponse, error) {
	params := &struct {
		Server string `json:"server,omitempty"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Token:  token,
	}
	return requestData[[]SteedRecordsResponse](ctx, c, "/steed/records", params)
}

// SteedRecord is an alias for SteedRecords.
func (c *Client) SteedRecord(ctx context.Context, server string, token string) (*[]SteedRecordsResponse, error) {
	return c.SteedRecords(ctx, server, token)
}
