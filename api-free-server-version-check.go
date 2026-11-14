package jx3api

import "context"

type ServerVersionCheckResponse struct {
	Num  int    `json:"num"`
	Size string `json:"size"`
	Ver  string `json:"ver"`
}

// ServerVersionCheck 查询游戏最新版本及更新包大小。
func (c *Client) ServerVersionCheck(ctx context.Context) (*ServerVersionCheckResponse, error) {
	return requestData[ServerVersionCheckResponse](ctx, c, "/server/version/check", nil)
}
