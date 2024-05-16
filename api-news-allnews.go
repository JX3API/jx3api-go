package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type NewsAllNewsResponse struct {
	ID    int    `json:"id"`
	Value int    `json:"value"`
	Type  string `json:"type"`
	Title string `json:"title"`
	Date  string `json:"date"`
	URL   string `json:"url"`
}

func (c *Client) NewsAllNews(ctx context.Context, limit int) (*[]NewsAllNewsResponse, error) {
	params := struct {
		Limit int `json:"limit"`
	}{
		Limit: limit,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("NewsAllNews: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/news/NewsAllNews", body)
	if err != nil {
		slog.Error("NewsAllNews: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("NewsAllNews: response body unmarshal error: ", err)
		return nil, err
	}

	data := new([]NewsAllNewsResponse)

	if resp.Msg != "success" {
		slog.Error("NewsAllNews: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("NewsAllNews: data unmarshal error: ", err)
		return nil, err
	}

	return data, nil
}
