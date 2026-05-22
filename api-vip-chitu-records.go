package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type ChituRecord struct {
	ID      int    `json:"id"`
	Server  string `json:"server"`
	MapName string `json:"map_name"`
	Horse   string `json:"horse"`
	Send    int    `json:"send"`
	Date    string `json:"date"`
}

func (c *Client) ChituRecords(ctx context.Context, token string) (*[]ChituRecord, error) {
	params := &struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("ChituRecords: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/chitu/records", body)
	if err != nil {
		slog.Error("ChituRecords: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("ChituRecords: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]ChituRecord)

	if resp.Msg != "success" {
		slog.Error("ChituRecords: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("ChituRecords: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
