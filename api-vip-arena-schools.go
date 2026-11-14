package jx3api

import "context"

type ArenaSchoolsResponse struct {
	ForceName string `json:"forceName"`
	LastWeek  int    `json:"lastWeek"`
	ThisWeek  int    `json:"thisWeek"`
}

// ArenaSchools 查询各门派的名剑大会周统计。
func (c *Client) ArenaSchools(ctx context.Context, mode int, ticket string, token string) (*[]ArenaSchoolsResponse, error) {
	params := &struct {
		Mode   int    `json:"mode"`
		Ticket string `json:"ticket,omitempty"`
		Token  string `json:"token,omitempty"`
	}{
		Mode:   mode,
		Ticket: ticket,
		Token:  token,
	}
	return requestData[[]ArenaSchoolsResponse](ctx, c, "/arena/schools", params)
}
