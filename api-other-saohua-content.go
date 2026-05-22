package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type SaohuaContentResponse struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

func (c *Client) SaohuaContent(ctx context.Context) (*SaohuaContentResponse, error) {
	params := &struct{}{}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("SaohuaContent: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/saohua/content", body)
	if err != nil {
		slog.Error("SaohuaContent: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("SaohuaContent: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(SaohuaContentResponse)

	if resp.Msg != "success" {
		slog.Error("SaohuaContent: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("SaohuaContent: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
