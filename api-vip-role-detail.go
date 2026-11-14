package jx3api

import "context"

type RoleDetailResponse struct {
	BodyId      int    `json:"bodyId"`
	BodyName    string `json:"bodyName"`
	CampId      int    `json:"campId"`
	CampName    string `json:"campName"`
	ForceId     int    `json:"forceId"`
	ForceName   string `json:"forceName"`
	GlobalId    string `json:"globalId"`
	RoleHistory *struct {
		TongNames []struct {
			Name   string `json:"name"`
			Server string `json:"server"`
			Time   int64  `json:"time"`
		} `json:"TongNames"`
		RoleNames []struct {
			Name   string `json:"name"`
			Server string `json:"server"`
			Time   int64  `json:"time"`
		} `json:"roleNames"`
	} `json:"roleHistory,omitempty"`
	RoleId     string `json:"roleId"`
	RoleName   string `json:"roleName"`
	ServerName string `json:"serverName"`
	TongId     int    `json:"tongId"`
	TongName   string `json:"tongName"`
	ZoneName   string `json:"zoneName"`
}

// RoleDetail 查询角色资料，可附带角色名和帮会历史。
func (c *Client) RoleDetail(ctx context.Context, server string, name string, history int, token string) (*RoleDetailResponse, error) {
	params := &struct {
		Server  string `json:"server"`
		Name    string `json:"name"`
		History int    `json:"history"`
		Token   string `json:"token,omitempty"`
	}{
		Server:  server,
		Name:    name,
		History: history,
		Token:   token,
	}
	return requestData[RoleDetailResponse](ctx, c, "/role/detail", params)
}
