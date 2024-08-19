package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type ExamAnswerResponse struct {
	ID          int    `json:"id"`
	Question    string `json:"question"`
	Answer      string `json:"answer"`
	Correctness int    `json:"correctness"`
	Index       int    `json:"index"`
	Pinyin      string `json:"pinyin"`
}

func (c *Client) ExamAnswer(ctx context.Context, subject string, limit int) (*[]ExamAnswerResponse, error) {
	params := &struct {
		Subject string `json:"subject"`
		Limit   int    `json:"limit"`
	}{
		Subject: subject,
		Limit:   limit,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("ExamAnswer: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/exam/answer", body)
	if err != nil {
		slog.Error("ExamAnswer: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("ExamAnswer: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]ExamAnswerResponse)

	if resp.Msg != "success" {
		slog.Error("ExamAnswer: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("ExamAnswer: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
