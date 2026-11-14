package jx3api

import "context"

type TrialSeasonsResponse struct {
	Bosses   []string `json:"bosses"`
	MaxFloor int      `json:"maxFloor"`
	Season   string   `json:"season"`
}

// TrialSeasons 查询当前试炼赛季及首领顺序。
func (c *Client) TrialSeasons(ctx context.Context, token string) (*[]TrialSeasonsResponse, error) {
	params := &struct {
		Token string `json:"token,omitempty"`
	}{
		Token: token,
	}
	return requestData[[]TrialSeasonsResponse](ctx, c, "/trial/seasons", params)
}
