package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type NextEventResponse struct {
	Zone   string `json:"zone"`
	Server string `json:"server"`
	Status int    `json:"status"`
	Time   int64  `json:"time"`
}

func (c *Client) NextEvent(ctx context.Context, server string, token string) (*[]NextEventResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Token  string `json:"token"`
	}{
		Server: server,
		Token:  token,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("NextEvent: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/active/next/event", body)
	if err != nil {
		slog.Error("NextEvent: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("NextEvent: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]NextEventResponse)

	if resp.Msg != "success" {
		slog.Error("NextEvent: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("NextEvent: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
