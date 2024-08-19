package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type FireworkStatisticalResponse struct {
	ID        int    `json:"id"`
	Zone      string `json:"zone"`
	Server    string `json:"server"`
	Name      string `json:"name"`
	MapName   string `json:"map_name"`
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Mode      int    `json:"mode"`
	Status    int    `json:"status"`
	Time      int64  `json:"time"`
}

func (c *Client) FireworkStatistical(ctx context.Context, server string, name string, limit int) (*[]FireworkStatisticalResponse, error) {
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
		slog.Error("FireworkStatistical: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/firework/statistical", body)
	if err != nil {
		slog.Error("FireworkStatistical: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("FireworkStatistical: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]FireworkStatisticalResponse)

	if resp.Msg != "success" {
		slog.Error("FireworkStatistical: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("FireworkStatistical: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
