package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type UnfinishedResponse struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Level int    `json:"level"`
}

func (c *Client) EventUnfinished(ctx context.Context, server string, name string, token string) (*[]UnfinishedResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Name   string `json:"name"`
		Token  string `json:"token"`
	}{
		Server: server,
		Name:   name,
		Token:  token,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("EventUnfinished: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/event/unfinished", body)
	if err != nil {
		slog.Error("EventUnfinished: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("EventUnfinished: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]UnfinishedResponse)

	if resp.Msg != "success" {
		slog.Error("EventUnfinished: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("EventUnfinished: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
