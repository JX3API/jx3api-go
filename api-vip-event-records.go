package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type EventRecordResponse struct {
	Zone    string `json:"zone"`
	Server  string `json:"server"`
	Name    string `json:"name"`
	Event   string `json:"event"`
	Level   int    `json:"level"`
	Status  int    `json:"status"`
	Time    int64  `json:"time"`
}

func (c *Client) EventRecords(ctx context.Context, server string, name string, token string) (*[]EventRecordResponse, error) {
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
		slog.Error("EventRecords: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/event/records", body)
	if err != nil {
		slog.Error("EventRecords: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("EventRecords: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]EventRecordResponse)

	if resp.Msg != "success" {
		slog.Error("EventRecords: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("EventRecords: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
