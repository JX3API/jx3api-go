package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type EventStatisticsResponse struct {
	ID      int    `json:"id"`
	Zone    string `json:"zone"`
	Server  string `json:"server"`
	Name    string `json:"name"`
	Event   string `json:"event"`
	Source  int    `json:"source"`
	Status  int    `json:"status"`
	Time    int64  `json:"time"`
}

func (c *Client) EventStatistics(ctx context.Context, server string, name string, limit int, token string) (*[]EventStatisticsResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Name   string `json:"name"`
		Limit  int    `json:"limit"`
		Token  string `json:"token"`
	}{
		Server: server,
		Name:   name,
		Limit:  limit,
		Token:  token,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("EventStatistics: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/event/statistics", body)
	if err != nil {
		slog.Error("EventStatistics: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("EventStatistics: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]EventStatisticsResponse)

	if resp.Msg != "success" {
		slog.Error("EventStatistics: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("EventStatistics: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
