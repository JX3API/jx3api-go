package jx3api

import "context"

type SkillReworkResponse struct {
	Id    int    `json:"id"`
	Time  string `json:"time"`
	Title string `json:"title"`
	Url   string `json:"url"`
}

// SkillRework 查询官方武学调整公告。
func (c *Client) SkillRework(ctx context.Context) (*[]SkillReworkResponse, error) {
	return requestData[[]SkillReworkResponse](ctx, c, "/skill/rework", nil)
}
