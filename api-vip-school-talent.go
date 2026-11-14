package jx3api

import "context"

type SchoolTalentResponse struct {
	Options []struct {
		Desc    string `json:"desc"`
		Icon    int    `json:"icon"`
		Id      int    `json:"id"`
		Level   int    `json:"level"`
		Name    string `json:"name"`
		Passive bool   `json:"passive"`
	} `json:"options"`
	PointId int `json:"pointId"`
	Slot    int `json:"slot"`
}

// SchoolTalent 查询心法的奇穴列表。
func (c *Client) SchoolTalent(ctx context.Context, kungfu string, token string) (*[]SchoolTalentResponse, error) {
	params := &struct {
		Kungfu string `json:"kungfu"`
		Token  string `json:"token,omitempty"`
	}{
		Kungfu: kungfu,
		Token:  token,
	}
	return requestData[[]SchoolTalentResponse](ctx, c, "/school/talent", params)
}
