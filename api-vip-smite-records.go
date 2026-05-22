package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type SmiteRecord struct {
	ID       int    `json:"id"`
	Zone     string `json:"zone"`
	Server   string `json:"server"`
	MapName  string `json:"map_name"`
	Time     int64  `json:"time"`
}

func (c *Client) SmiteRecords(ctx context.Context, token string) (*[]SmiteRecord, error) {
	params := &struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("SmiteRecords: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/smite/records", body)
	if err != nil {
		slog.Error("SmiteRecords: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("SmiteRecords: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]SmiteRecord)

	if resp.Msg != "success" {
		slog.Error("SmiteRecords: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("SmiteRecords: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
