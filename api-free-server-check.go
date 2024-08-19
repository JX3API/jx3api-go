package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type ServerCheckResponse struct {
	ID     int    `json:"id"`
	Zone   string `json:"zone"`
	Server string `json:"server"`
	Status int    `json:"status"`
	Time   int64  `json:"time"`
}

func (c *Client) ServerCheck(ctx context.Context, server string) (*ServerCheckResponse, error) {
	params := &struct {
		Server string `json:"server"`
	}{
		Server: server,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("ServerCheck: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/server/check", body)
	if err != nil {
		slog.Error("ServerCheck: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("ServerCheck: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(ServerCheckResponse)

	if resp.Msg != "success" {
		slog.Error("ServerCheck: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("ServerCheck: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
