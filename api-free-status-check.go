package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type StatusCheckResponse struct {
	Zone   string `json:"zone"`
	Server string `json:"server"`
	Status string `json:"status"`
}

func (c *Client) StatusCheck(ctx context.Context, typ int, server string) (*StatusCheckResponse, error) {
	params := &struct {
		Type   int    `json:"type"`
		Server string `json:"server"`
	}{
		Type:   typ,
		Server: server,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("StatusCheck: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/status/check", body)
	if err != nil {
		slog.Error("StatusCheck: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("StatusCheck: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(StatusCheckResponse)

	if resp.Msg != "success" {
		slog.Error("StatusCheck: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("StatusCheck: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
