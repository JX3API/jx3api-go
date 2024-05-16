package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type SeverCheckResponse struct {
	ID     int    `json:"id"`
	Zone   string `json:"zone"`
	Server string `json:"server"`
	Status int    `json:"status"`
	Time   int    `json:"time"`
}

func (c *Client) SeverCheck(ctx context.Context, server string) (*SeverCheckResponse, error) {
	params := struct {
		Server string `json:"server"`
	}{
		Server: server,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("SeverCheck: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/server/check", body)
	if err != nil {
		slog.Error("SeverCheck: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("SeverCheck: response body unmarshal error: ", err)
		return nil, err
	}

	data := new(SeverCheckResponse)

	if resp.Msg != "success" {
		slog.Error("SeverCheck: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("SeverCheck: data unmarshal error: ", err)
		return nil, err
	}

	return data, nil
}
