package jx3api

import "context"

type ExamSearchResponse struct {
	Answer      string `json:"answer"`
	Correctness *int   `json:"correctness"`
	Index       int    `json:"index"`
	Question    string `json:"question"`
}

type ExamAnswerResponse = ExamSearchResponse

// ExamSearch 搜索科举试题及答案。
func (c *Client) ExamSearch(ctx context.Context, subject string, limit int) (*[]ExamSearchResponse, error) {
	params := &struct {
		Subject string `json:"subject"`
		Limit   int    `json:"limit"`
	}{
		Subject: subject,
		Limit:   limit,
	}
	return requestData[[]ExamSearchResponse](ctx, c, "/exam/search", params)
}

// ExamAnswer is an alias for ExamSearch.
func (c *Client) ExamAnswer(ctx context.Context, subject string, limit int) (*[]ExamSearchResponse, error) {
	return c.ExamSearch(ctx, subject, limit)
}
