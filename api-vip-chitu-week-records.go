package jx3api

import "context"

type ChituWeekRecordsResponse = ChituRecord

// ChituWeekRecords 查询本周赤兔预告与出世记录，每周二 07:00（北京时间）重置。
func (c *Client) ChituWeekRecords(ctx context.Context, token string) (*[]ChituWeekRecordsResponse, error) {
	params := &struct {
		Token string `json:"token,omitempty"`
	}{
		Token: token,
	}
	return requestData[[]ChituWeekRecordsResponse](ctx, c, "/chitu/week/records", params)
}
