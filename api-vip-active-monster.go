package jx3api

import "context"

type MonsterWeeklyResponse struct {
	Boss string `json:"boss"`
	End  int64  `json:"end"`
	List []struct {
		Data struct {
			Desc string   `json:"desc"`
			List []string `json:"list"`
			Name string   `json:"name"`
		} `json:"data"`
		Index int      `json:"index"`
		Name  string   `json:"name"`
		Skill []string `json:"skill"`
	} `json:"list"`
	Start int64  `json:"start"`
	Week  string `json:"week"`
}

type MonsterResponse = MonsterWeeklyResponse

// MonsterWeekly 查询本周百战首领、技能及效果。
func (c *Client) MonsterWeekly(ctx context.Context, token string) (*MonsterWeeklyResponse, error) {
	params := &struct {
		Token string `json:"token,omitempty"`
	}{
		Token: token,
	}
	return requestData[MonsterWeeklyResponse](ctx, c, "/monster/weekly", params)
}

// ActiveMonster is an alias for MonsterWeekly.
func (c *Client) ActiveMonster(ctx context.Context, token string) (*MonsterWeeklyResponse, error) {
	return c.MonsterWeekly(ctx, token)
}
