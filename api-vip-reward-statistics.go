package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type RewardStatisticsResponse struct {
	ID        int    `json:"id"`
	Zone      string `json:"zone"`
	Server    string `json:"server"`
	MapName   string `json:"map_name"`
	RoleName  string `json:"role_name"`
	ItemName  string `json:"item_name"`
	ItemAmount string `json:"item_amount"`
	Time      int64  `json:"time"`
}

func (c *Client) RewardStatistics(ctx context.Context, server string, name string, limit int, token string) (*[]RewardStatisticsResponse, error) {
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
		slog.Error("RewardStatistics: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/reward/statistics", body)
	if err != nil {
		slog.Error("RewardStatistics: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("RewardStatistics: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]RewardStatisticsResponse)

	if resp.Msg != "success" {
		slog.Error("RewardStatistics: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("RewardStatistics: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
