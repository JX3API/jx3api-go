package jx3api

import "context"

type MentorSearchResponse struct {
	BodyId          int    `json:"bodyId"`
	BodyName        string `json:"bodyName"`
	CampId          int    `json:"campId"`
	CampName        string `json:"campName"`
	CommentText     string `json:"commentText"`
	ForceId         int    `json:"forceId"`
	ForceName       string `json:"forceName"`
	MapId           int    `json:"mapId"`
	MiniAvatorId    int    `json:"miniAvatorId"`
	RoleId          int    `json:"roleId"`
	RoleLevel       int    `json:"roleLevel"`
	RoleName        string `json:"roleName"`
	TongId          int    `json:"tongId"`
	TongMasterName  string `json:"tongMasterName"`
	TongMemberCount int    `json:"tongMemberCount"`
	TongName        string `json:"tongName"`
}

// MentorSearch 查询服务器的师徒招募信息。
func (c *Client) MentorSearch(ctx context.Context, typ int, server string, keyword string, limit int, token string) (*[]MentorSearchResponse, error) {
	params := &struct {
		Label   int    `json:"label"`
		Server  string `json:"server"`
		Keyword string `json:"keyword,omitempty"`
		Limit   int    `json:"limit"`
		Token   string `json:"token,omitempty"`
	}{
		Label:   typ,
		Server:  server,
		Keyword: keyword,
		Limit:   limit,
		Token:   token,
	}
	return requestData[[]MentorSearchResponse](ctx, c, "/mentor/search", params)
}
