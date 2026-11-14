package jx3api

import "context"

type SkillCalculateResponse struct {
	Duration          float64 `json:"duration"`
	HastePercent      float64 `json:"hastePercent"`
	HastePercentLimit float64 `json:"hastePercentLimit"`
	Level             int     `json:"level"`
	NowFrame          int     `json:"nowFrame"`
	SurplusNowFrame   int     `json:"surplusNowFrame"`
	UNowFrame         int     `json:"uNowFrame"`
}

// SkillCalculate 计算技能在各急速档位的读条时长。
func (c *Client) SkillCalculate(ctx context.Context, cooldown float64, token string) (*[]SkillCalculateResponse, error) {
	params := &struct {
		Cooldown float64 `json:"cooldown"`
		Token    string  `json:"token,omitempty"`
	}{
		Cooldown: cooldown,
		Token:    token,
	}
	return requestData[[]SkillCalculateResponse](ctx, c, "/skill/calculate", params)
}
