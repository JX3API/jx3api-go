package jx3api

import "context"

type ChatRecordsResponse struct {
	List []struct {
		Channel  string `json:"channel"`
		GlobalId string `json:"globalId"`
		Message  string `json:"message"`
		RoleId   string `json:"roleId"`
		RoleName string `json:"roleName"`
		Server   string `json:"server"`
		Time     int64  `json:"time"`
		Zone     string `json:"zone"`
	} `json:"list"`
	Total int `json:"total"`
}

// ChatRecords 分页查询角色聊天记录。
func (c *Client) ChatRecords(ctx context.Context, server string, name string, limit int, page int, token string) (*ChatRecordsResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Name   string `json:"name"`
		Limit  int    `json:"limit"`
		Page   int    `json:"page"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Name:   name,
		Limit:  limit,
		Page:   page,
		Token:  token,
	}
	return requestData[ChatRecordsResponse](ctx, c, "/chat/records", params)
}
