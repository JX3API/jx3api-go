package jx3api

import (
	"context"
	"encoding/json"
)

type FraudDetailResponse struct {
	Author    string          `json:"author"`
	Commented bool            `json:"commented"`
	Content   string          `json:"content"`
	Created   int             `json:"created"`
	Pid       int64           `json:"pid"`
	Status    int             `json:"status"`
	Tid       int64           `json:"tid"`
	Tieba     string          `json:"tieba"`
	Title     string          `json:"title"`
	Uid       json.RawMessage `json:"uid"`
}

// FraudDetail 查询角色 UID 对应的骗子记录。
func (c *Client) FraudDetail(ctx context.Context, uid int, server string, token string) (*[]FraudDetailResponse, error) {
	params := &struct {
		Uid    int    `json:"uid"`
		Server string `json:"server,omitempty"`
		Token  string `json:"token,omitempty"`
	}{
		Uid:    uid,
		Server: server,
		Token:  token,
	}
	return requestData[[]FraudDetailResponse](ctx, c, "/fraud/detail", params)
}
