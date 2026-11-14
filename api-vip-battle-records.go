package jx3api

import "context"

type BattleRecordsResponse struct {
	AcceptingName   string  `json:"acceptingName"`
	DeclaringName   string  `json:"declaringName"`
	DurationSeconds int     `json:"durationSeconds"`
	EndTime         *int64  `json:"endTime"`
	StartTime       int64   `json:"startTime"`
	Type            string  `json:"type"`
	WinnerName      *string `json:"winnerName,omitempty"`
	WinnerScore     *string `json:"winnerScore,omitempty"`
}

// BattleRecords 查询服务器最近 30 场帮战。
func (c *Client) BattleRecords(ctx context.Context, server string) (*[]BattleRecordsResponse, error) {
	params := &struct {
		Server string `json:"server"`
	}{
		Server: server,
	}
	return requestData[[]BattleRecordsResponse](ctx, c, "/battle/records", params)
}
