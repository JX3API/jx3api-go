package jx3api

import "context"

type CardRandomResponse = CardCachedResponse

// CardRandom 随机获取一张角色名片。
func (c *Client) CardRandom(ctx context.Context, server string, body string, force string, token string) (*CardRandomResponse, error) {
	params := &struct {
		Server string `json:"server,omitempty"`
		Body   string `json:"body,omitempty"`
		Force  string `json:"force,omitempty"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Body:   body,
		Force:  force,
		Token:  token,
	}
	return requestData[CardRandomResponse](ctx, c, "/card/random", params)
}
