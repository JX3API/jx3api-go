package jx3api

import "context"

type ChoiceAnswerResponse struct {
	Answer  string `json:"answer"`
	Hearten string `json:"hearten"`
}

// ChoiceAnswer 随机获取一条答案和一句鼓励语。
func (c *Client) ChoiceAnswer(ctx context.Context) (*ChoiceAnswerResponse, error) {
	return requestData[ChoiceAnswerResponse](ctx, c, "/choice/answer", nil)
}
