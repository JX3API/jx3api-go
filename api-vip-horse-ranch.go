package jx3api

import "context"

type RanchRecordsResponse struct {
	Data   map[string][]string `json:"data"`
	Note   string              `json:"note"`
	Server string              `json:"server"`
	Zone   string              `json:"zone"`
}

// RanchRecords 查询马场预告及本周期赤兔、的卢出世情况。
func (c *Client) RanchRecords(ctx context.Context, server string, expired int, token string) (*RanchRecordsResponse, error) {
	params := &struct {
		Server  string `json:"server"`
		Expired int    `json:"expired"`
		Token   string `json:"token,omitempty"`
	}{
		Server:  server,
		Expired: expired,
		Token:   token,
	}
	return requestData[RanchRecordsResponse](ctx, c, "/ranch/records", params)
}
