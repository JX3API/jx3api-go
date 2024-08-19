package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type AntiviceResponse struct {
	ID      int    `json:"id"`
	Zone    string `json:"zone"`
	Server  string `json:"server"`
	MapName string `json:"map_name"`
	Time    int    `json:"time"`
}

func (c *Client) ServerAntivice(ctx context.Context, server string) (*[]AntiviceResponse, error) {
	params := &struct {
		Server string `json:"server"`
	}{
		Server: server,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("ServerAntivice: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/server/antivice", body)
	if err != nil {
		slog.Error("ServerAntivice: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("ServerAntivice: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]AntiviceResponse)

	if resp.Msg != "success" {
		slog.Error("ServerAntivice: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("ServerAntivice: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
