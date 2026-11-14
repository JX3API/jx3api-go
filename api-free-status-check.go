package jx3api

import "context"

type ServerStatus struct {
	Zone   string `json:"zone"`
	Server string `json:"server"`
	// Status is an integer for type=1 and a string for official online status.
	Status   any   `json:"status"`
	Lasttime int64 `json:"lasttime,omitempty"`
	Shuttime int64 `json:"shuttime,omitempty"`
}
type ServerStatusCheckResponse []ServerStatus

func (r *ServerStatusCheckResponse) UnmarshalJSON(raw []byte) error {
	return unmarshalRecords(raw, (*[]ServerStatus)(r))
}

type StatusCheckResponse = ServerStatusCheckResponse

// ServerStatusCheck 查询服务器的开服或在线状态。
func (c *Client) ServerStatusCheck(ctx context.Context, typ int, server string) (*ServerStatusCheckResponse, error) {
	params := &struct {
		Type   int    `json:"type"`
		Server string `json:"server,omitempty"`
	}{
		Type:   typ,
		Server: server,
	}
	return requestData[ServerStatusCheckResponse](ctx, c, "/server/status/check", params)
}

// StatusCheck is an alias for ServerStatusCheck.
func (c *Client) StatusCheck(ctx context.Context, typ int, server string) (*ServerStatusCheckResponse, error) {
	return c.ServerStatusCheck(ctx, typ, server)
}
