package jx3api

import "context"

type CalcdRecordsResponse struct {
	EquipList []struct {
		CacheTime  int64  `json:"cacheTime"`
		EquipType  string `json:"equipType"`
		KungfuName string `json:"kungfuName"`
		SeasonName string `json:"seasonName"`
	} `json:"equipList"`
	GlobalId   string `json:"globalId"`
	RoleId     string `json:"roleId"`
	RoleName   string `json:"roleName"`
	ServerName string `json:"serverName"`
	ZoneName   string `json:"zoneName"`
}

// CalcdRecords 查询角色已收录的配装与赛季。
func (c *Client) CalcdRecords(ctx context.Context, server string, name string, token string) (*CalcdRecordsResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Name   string `json:"name"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Name:   name,
		Token:  token,
	}
	return requestData[CalcdRecordsResponse](ctx, c, "/calcd/records", params)
}
