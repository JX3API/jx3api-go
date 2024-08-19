package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type NewsAllnewsResponse struct {
	ID    int    `json:"id"`
	Token int    `json:"token"`
	Class string `json:"class"`
	Title string `json:"title"`
	Date  string `json:"date"`
	URL   string `json:"url"`
}

func (c *Client) NewsAllnews(ctx context.Context, limit int) (*[]NewsAllnewsResponse, error) {
	params := &struct {
		Limit int `json:"limit"`
	}{
		Limit: limit,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("NewsAllnews: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/news/allnews", body)
	if err != nil {
		slog.Error("NewsAllnews: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("NewsAllnews: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]NewsAllnewsResponse)

	if resp.Msg != "success" {
		slog.Error("NewsAllnews: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("NewsAllnews: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
