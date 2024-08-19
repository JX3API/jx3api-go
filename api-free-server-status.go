package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type ServerStatusResponse struct {
	Zone   string `json:"zone"`
	Server string `json:"server"`
	Status string `json:"status"`
}

func (c *Client) ServerStatus(ctx context.Context, server string) (*ServerStatusResponse, error) {
	params := &struct {
		Server string `json:"server"`
	}{
		Server: server,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("ServerStatus: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/server/status", body)
	if err != nil {
		slog.Error("ServerStatus: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("ServerStatus: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(ServerStatusResponse)

	if resp.Msg != "success" {
		slog.Error("ServerStatus: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("ServerStatus: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
