package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type SaohuaRandomResponse struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

func (c *Client) SaohuaRandom(ctx context.Context) (*SaohuaRandomResponse, error) {
	params := &struct{}{}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("SaohuaRandom: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/saohua/random", body)
	if err != nil {
		slog.Error("SaohuaRandom: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("SaohuaRandom: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(SaohuaRandomResponse)

	if resp.Msg != "success" {
		slog.Error("SaohuaRandom: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("SaohuaRandom: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
