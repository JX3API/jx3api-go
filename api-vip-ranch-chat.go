package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
)

type RanchChatResponse map[string][]string

func (r *RanchChatResponse) UnmarshalJSON(raw []byte) error {
	if bytes.Equal(bytes.TrimSpace(raw), []byte("[]")) {
		*r = nil
		return nil
	}
	type plain RanchChatResponse
	return json.Unmarshal(raw, (*plain)(r))
}

// RanchChat 查询各地图 NPC 播报的马匹预计出世时间。
func (c *Client) RanchChat(ctx context.Context, server string, token string) (*RanchChatResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Token:  token,
	}
	return requestData[RanchChatResponse](ctx, c, "/ranch/chat", params)
}
