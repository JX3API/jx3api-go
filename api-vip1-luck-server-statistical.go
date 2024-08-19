package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type LuckServerStatisticalResponse struct {
	ID     int64  `json:"id"`
	Zone   string `json:"zone"`
	Server string `json:"server"`
	Name   string `json:"name"`
	Event  string `json:"event"`
	Status int    `json:"status"`
	Time   int64  `json:"time"`
}

func (c *Client) LuckServerStatistical(ctx context.Context, name string, limit int) (*[]LuckServerStatisticalResponse, error) {
	params := &struct {
		Name  string `json:"name"`
		Limit int    `json:"limit"`
	}{
		Name:  name,
		Limit: limit,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("ServerStatistical: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/luck/server/statistical", body)
	if err != nil {
		slog.Error("ServerStatistical: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("ServerStatistical: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]LuckServerStatisticalResponse)

	if resp.Msg != "success" {
		slog.Error("ServerStatistical: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("ServerStatistical: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
