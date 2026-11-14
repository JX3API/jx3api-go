package jx3api

import "context"

type ActiveCelebsResponse struct {
	Desc  string `json:"desc"`
	Icon  string `json:"icon"`
	Map   string `json:"map"`
	Site  string `json:"site"`
	Stage string `json:"stage"`
	Time  string `json:"time"`
}

// ActiveCelebs 查询指定地图活动当前及后续的 10 条安排，时间为北京时间。
func (c *Client) ActiveCelebs(ctx context.Context, name string) (*[]ActiveCelebsResponse, error) {
	params := &struct {
		Name string `json:"name"`
	}{
		Name: name,
	}
	return requestData[[]ActiveCelebsResponse](ctx, c, "/active/celebs", params)
}
