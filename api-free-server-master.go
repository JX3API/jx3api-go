package jx3api

import "context"

type ServerSearchResponse struct {
	Server string `json:"server"`
}

type ServerMasterResponse = ServerSearchResponse

// ServerSearch 查询名称对应的主服务器。
func (c *Client) ServerSearch(ctx context.Context, server string) (*ServerSearchResponse, error) {
	params := &struct {
		Server string `json:"server"`
	}{
		Server: server,
	}
	return requestData[ServerSearchResponse](ctx, c, "/server/search", params)
}

// ServerMaster is an alias for ServerSearch.
func (c *Client) ServerMaster(ctx context.Context, name string) (*ServerSearchResponse, error) {
	return c.ServerSearch(ctx, name)
}
