package jx3api

import "context"

type MechDecryptResponse struct {
	IntervalTime string `json:"intervalTime"`
	NextNode     string `json:"nextNode"`
	NextResult   string `json:"nextResult"`
	NowNode      string `json:"nowNode"`
	NowResult    string `json:"nowResult"`
	NowTime      string `json:"nowTime"`
}

type MechCalculatorResponse = MechDecryptResponse

// MechDecrypt 查询当前秘境方位及下次变化的倒计时。
func (c *Client) MechDecrypt(ctx context.Context, token string) (*MechDecryptResponse, error) {
	params := &struct {
		Token string `json:"token,omitempty"`
	}{
		Token: token,
	}
	return requestData[MechDecryptResponse](ctx, c, "/mech/decrypt", params)
}

// MechCalculator is an alias for MechDecrypt.
func (c *Client) MechCalculator(ctx context.Context) (*MechDecryptResponse, error) {
	return c.MechDecrypt(ctx, "")
}
