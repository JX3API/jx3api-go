package jx3api

import "context"

type CardCachedResponse struct {
	CacheTime    int64  `json:"cacheTime"`
	FileName     string `json:"fileName"`
	GlobalRoleId string `json:"globalRoleId"`
	RoleId       string `json:"roleId"`
	RoleName     string `json:"roleName"`
	ServerName   string `json:"serverName"`
	ShowActive   bool   `json:"showActive"`
	ShowAvatar   string `json:"showAvatar"`
	ShowHash     string `json:"showHash"`
	ShowIndex    int    `json:"showIndex"`
	ShowLike     int    `json:"showLike"`
	ZoneName     string `json:"zoneName"`
}

// CardCached 查询已缓存的角色名片。
func (c *Client) CardCached(ctx context.Context, server string, name string, token string) (*CardCachedResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Name   string `json:"name"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Name:   name,
		Token:  token,
	}
	return requestData[CardCachedResponse](ctx, c, "/card/cached", params)
}
