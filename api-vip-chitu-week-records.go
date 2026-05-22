package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type ChituWeekRecord struct {
	Server  string `json:"server"`
	MapName string `json:"map_name"`
	Horse   string `json:"horse"`
	Date    string `json:"date"`
}

func (c *Client) ChituWeekRecords(ctx context.Context, token string) (*[]ChituWeekRecord, error) {
	params := &struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("ChituWeekRecords: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/chitu/week/records", body)
	if err != nil {
		slog.Error("ChituWeekRecords: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("ChituWeekRecords: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]ChituWeekRecord)

	if resp.Msg != "success" {
		slog.Error("ChituWeekRecords: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("ChituWeekRecords: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
