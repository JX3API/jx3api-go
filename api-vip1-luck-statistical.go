package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type LuckStatisticalResponse struct {
	ID     int64  `json:"id"`
	Zone   string `json:"zone"`
	Server string `json:"server"`
	Name   string `json:"name"`
	Event  string `json:"event"`
	Status int    `json:"status"`
	Time   int64  `json:"time"`
}

func (c *Client) LuckStatistical(ctx context.Context, server string, name string, limit int) (*[]LuckStatisticalResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Name   string `json:"name"`
		Limit  int    `json:"limit"`
	}{
		Server: server,
		Name:   name,
		Limit:  limit,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("LuckStatistical: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/luck/statistical", body)
	if err != nil {
		slog.Error("LuckStatistical: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("LuckStatistical: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]LuckStatisticalResponse)

	if resp.Msg != "success" {
		slog.Error("LuckStatistical: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("LuckStatistical: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
