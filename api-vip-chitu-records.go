package jx3api

import "context"

type ChituRecord struct {
	Server  string  `json:"server"`
	MapName string  `json:"mapName"`
	Horse   string  `json:"horse"`
	Date    string  `json:"date"`
	Time    *int64  `json:"time"`
	Window  *string `json:"window"`
	Status  string  `json:"status"`
}
type ChituRecordsResponse = ChituRecord

// ChituRecords 查询今日赤兔预告与出世记录。
func (c *Client) ChituRecords(ctx context.Context, token string) (*[]ChituRecordsResponse, error) {
	params := &struct {
		Token string `json:"token,omitempty"`
	}{
		Token: token,
	}
	return requestData[[]ChituRecordsResponse](ctx, c, "/chitu/records", params)
}
