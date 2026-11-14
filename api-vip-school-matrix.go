package jx3api

import "context"

type SchoolMatrixResponse struct {
	Desc   string `json:"desc"`
	Icon   int    `json:"icon"`
	Id     int    `json:"id"`
	Levels []struct {
		Desc  string `json:"desc"`
		Level int    `json:"level"`
	} `json:"levels"`
	Name string `json:"name"`
}

// SchoolMatrix 查询心法的阵法列表。
func (c *Client) SchoolMatrix(ctx context.Context, kungfu string, token string) (*[]SchoolMatrixResponse, error) {
	params := &struct {
		Kungfu string `json:"kungfu"`
		Token  string `json:"token,omitempty"`
	}{
		Kungfu: kungfu,
		Token:  token,
	}
	return requestData[[]SchoolMatrixResponse](ctx, c, "/school/matrix", params)
}
