package jx3api

import "context"

type SchoolSearchResponse struct {
	CommentCount    int    `json:"commentCount"`
	Content         string `json:"content"`
	EquipShareId    string `json:"equipShareId"`
	Force           string `json:"force"`
	Id              string `json:"id"`
	KungfuId        string `json:"kungfuId"`
	KungfuName      string `json:"kungfuName"`
	LikeCount       int    `json:"likeCount"`
	PublishTimeUnix int    `json:"publishTimeUnix"`
	ReplyTimeUnix   int    `json:"replyTimeUnix"`
	RoleInfo        struct {
		RoleName   string `json:"roleName"`
		SchoolName string `json:"schoolName"`
		ServerName string `json:"serverName"`
		ZoneName   string `json:"zoneName"`
	} `json:"roleInfo"`
	Score      int      `json:"score"`
	ShareCount int      `json:"shareCount"`
	Tags       []string `json:"tags"`
	Title      string   `json:"title"`
	TuilanInfo struct {
		AvatarUrl string `json:"avatarUrl"`
		Nickname  string `json:"nickname"`
		PersonId  string `json:"personId"`
	} `json:"tuilanInfo"`
}

// SchoolSearch 查询指定心法的推荐配装。
func (c *Client) SchoolSearch(ctx context.Context, name string, mode string, ticket string, token string) (*[]SchoolSearchResponse, error) {
	params := &struct {
		Name   string `json:"name"`
		Mode   string `json:"mode,omitempty"`
		Ticket string `json:"ticket,omitempty"`
		Token  string `json:"token,omitempty"`
	}{
		Name:   name,
		Mode:   mode,
		Ticket: ticket,
		Token:  token,
	}
	return requestData[[]SchoolSearchResponse](ctx, c, "/school/search", params)
}
