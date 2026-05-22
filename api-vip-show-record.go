package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type ShowRecordResponse struct {
	ID        int    `json:"id"`
	Zone      string `json:"zone"`
	Server    string `json:"server"`
	MapName   string `json:"map_name"`
	Sender    string `json:"sender"`
	Receiver  string `json:"receiver"`
	Firework  string `json:"firework"`
	Matched   int    `json:"matched"`
	Status    int    `json:"status"`
	Time      int64  `json:"time"`
}

func (c *Client) ShowRecord(ctx context.Context, server string, name string, token string) (*[]ShowRecordResponse, error) {
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
		slog.Error("ShowRecord: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/show/records", body)
	if err != nil {
		slog.Error("FireworkRecord: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("ShowRecord: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]ShowRecordResponse)

	if resp.Msg != "success" {
		slog.Error("ShowRecord: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("ShowRecord: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
