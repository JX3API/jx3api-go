package jx3api

import "context"

type ArenaAwesomeResponse struct {
	AvatarUrl  string `json:"avatarUrl"`
	ForceName  string `json:"forceName"`
	MatchScore string `json:"matchScore"`
	RankNum    string `json:"rankNum"`
	RoleName   string `json:"roleName"`
	UpNum      int    `json:"upNum"`
	WinRate    int    `json:"winRate"`
}

// ArenaAwesome 查询名剑大会周排行榜。
func (c *Client) ArenaAwesome(ctx context.Context, mode int, limit int, ticket string, token string) (*[]ArenaAwesomeResponse, error) {
	params := &struct {
		Mode   int    `json:"mode"`
		Limit  int    `json:"limit"`
		Ticket string `json:"ticket,omitempty"`
		Token  string `json:"token,omitempty"`
	}{
		Mode:   mode,
		Limit:  limit,
		Ticket: ticket,
		Token:  token,
	}
	return requestData[[]ArenaAwesomeResponse](ctx, c, "/arena/awesome", params)
}
