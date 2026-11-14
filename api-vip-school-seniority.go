package jx3api

import "context"

type SchoolSeniorityResponse struct {
	AvatarUrl  string `json:"avatarUrl"`
	ForceIcon  string `json:"forceIcon"`
	ForceId    int    `json:"forceId"`
	ForceName  string `json:"forceName"`
	RoleId     int    `json:"roleId"`
	RoleName   string `json:"roleName"`
	Seniority  int    `json:"seniority"`
	ServerName string `json:"serverName"`
	ZoneName   string `json:"zoneName"`
}

// SchoolSeniority 查询全服或指定服务器、门派的资历前 100 名。
func (c *Client) SchoolSeniority(ctx context.Context, server string, school string, ticket string, token string) (*[]SchoolSeniorityResponse, error) {
	params := &struct {
		Server string `json:"server,omitempty"`
		School string `json:"school,omitempty"`
		Ticket string `json:"ticket,omitempty"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		School: school,
		Ticket: ticket,
		Token:  token,
	}
	return requestData[[]SchoolSeniorityResponse](ctx, c, "/school/seniority", params)
}
