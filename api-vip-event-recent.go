package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type EventRecentResponse struct {
	ID      int    `json:"id"`
	Zone    string `json:"zone"`
	Server  string `json:"server"`
	Name    string `json:"name"`
	Event   string `json:"event"`
	Source  int    `json:"source"`
	Status  int    `json:"status"`
	Time    int64  `json:"time"`
}

func (c *Client) EventRecent(ctx context.Context, server string, token string) (*[]EventRecentResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Token  string `json:"token"`
	}{
		Server: server,
		Token:  token,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("EventRecent: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/event/recent", body)
	if err != nil {
		slog.Error("EventRecent: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("EventRecent: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]EventRecentResponse)

	if resp.Msg != "success" {
		slog.Error("EventRecent: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("EventRecent: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
