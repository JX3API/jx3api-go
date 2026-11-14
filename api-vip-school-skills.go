package jx3api

import "context"

type SchoolSkillsResponse struct {
	Desc        string  `json:"desc"`
	Distance    int     `json:"distance"`
	Group       string  `json:"group"`
	HelpDesc    string  `json:"helpDesc"`
	Icon        int     `json:"icon"`
	Id          int     `json:"id"`
	Interval    float64 `json:"interval"`
	Name        string  `json:"name"`
	Passive     bool    `json:"passive"`
	SpecialDesc string  `json:"specialDesc"`
}

// SchoolSkills 查询心法的技能列表。
func (c *Client) SchoolSkills(ctx context.Context, kungfu string, token string) (*[]SchoolSkillsResponse, error) {
	params := &struct {
		Kungfu string `json:"kungfu"`
		Token  string `json:"token,omitempty"`
	}{
		Kungfu: kungfu,
		Token:  token,
	}
	return requestData[[]SchoolSkillsResponse](ctx, c, "/school/skills", params)
}
