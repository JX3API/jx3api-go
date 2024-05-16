package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type SeverStatusResponse struct {
	Zone   string `json:"zone"`
	Server string `json:"server"`
	Status int    `json:"status"`
}

func (c *Client) SeverStatus(ctx context.Context, server string) (*SeverStatusResponse, error) {
	params := struct {
		Server string `json:"server"`
	}{
		Server: server,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("SeverStatus: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/server/status", body)
	if err != nil {
		slog.Error("SeverStatus: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("SeverStatus: response body unmarshal error: ", err)
		return nil, err
	}

	data := new(SeverStatusResponse)

	if resp.Msg != "success" {
		slog.Error("SeverStatus: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("SeverStatus: data unmarshal error: ", err)
		return nil, err
	}

	return data, nil
}
