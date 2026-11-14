package jx3api

import "context"

type HomeFlowerResponse map[string][]struct {
	Color string   `json:"color"`
	Line  []string `json:"line"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
}

// HomeFlower 查询各家园地图的鲜花收购信息。
func (c *Client) HomeFlower(ctx context.Context, server string, name string, mapName string) (*HomeFlowerResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Name   string `json:"name,omitempty"`
		Map    string `json:"map,omitempty"`
	}{
		Server: server,
		Name:   name,
		Map:    mapName,
	}
	return requestData[HomeFlowerResponse](ctx, c, "/home/flower", params)
}
