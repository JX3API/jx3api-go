package jx3api

import "context"

type TrialBossesResponse struct {
	Bosses []struct {
		Boss       string  `json:"boss"`
		ClearDps   *int    `json:"clearDps"`
		Group      string  `json:"group"`
		Hp         int     `json:"hp"`
		Note       *string `json:"note"`
		PerfectDps *int    `json:"perfectDps"`
	} `json:"bosses"`
	Floor    int     `json:"floor"`
	HpGrowth float64 `json:"hpGrowth"`
	Season   string  `json:"season"`
}

// TrialBosses 查询当前赛季指定层数的首领血量与通关秒伤。
func (c *Client) TrialBosses(ctx context.Context, floor int, token string) (*TrialBossesResponse, error) {
	params := &struct {
		Floor int    `json:"floor"`
		Token string `json:"token,omitempty"`
	}{
		Floor: floor,
		Token: token,
	}
	return requestData[TrialBossesResponse](ctx, c, "/trial/bosses", params)
}
