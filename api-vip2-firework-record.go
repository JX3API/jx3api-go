package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type FireworkRecordResponse struct {
	ID        int    `json:"id"`
	Zone      string `json:"zone"`
	Server    string `json:"server"`
	Name      string `json:"name"`
	MapName   string `json:"map_name"`
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Status    int    `json:"status"`
	Time      int64  `json:"time"`
}

func (c *Client) FireworkRecord(ctx context.Context, server string, name string) (*[]FireworkRecordResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Name   string `json:"name"`
	}{
		Server: server,
		Name:   name,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("FireworkRecord: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/firework/record", body)
	if err != nil {
		slog.Error("FireworkRecord: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("FireworkRecord: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]FireworkRecordResponse)

	if resp.Msg != "success" {
		slog.Error("FireworkRecord: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("FireworkRecord: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
