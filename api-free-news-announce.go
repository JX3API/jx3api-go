package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type NewsAnnounceResponse struct {
	ID    int    `json:"id"`
	Token int    `json:"token"`
	Class string `json:"class"`
	Title string `json:"title"`
	Date  string `json:"date"`
	URL   string `json:"url"`
}

func (c *Client) NewsAnnounce(ctx context.Context, limit int) (*[]NewsAnnounceResponse, error) {
	params := &struct {
		Limit int `json:"limit"`
	}{
		Limit: limit,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("NewsAnnounce: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/news/announce", body)
	if err != nil {
		slog.Error("NewsAnnounce: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("NewsAnnounce: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]NewsAnnounceResponse)

	if resp.Msg != "success" {
		slog.Error("NewsAnnounce: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("NewsAnnounce: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
