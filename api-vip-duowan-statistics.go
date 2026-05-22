package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type DuowanStatisticsData struct {
	Sid      int    `json:"sid"`
	LogoUrl  string `json:"logoUrl"`
	Users    int    `json:"users"`
	Snick    string `json:"snick"`
	Limit    int    `json:"limit"`
	Logo     int    `json:"logo"`
	Asid     int    `json:"asid"`
	Esid     string `json:"esid"`
	CampName string `json:"campName"`
}

type DuowanStatisticsServer struct {
	Server string                 `json:"server"`
	Data   []DuowanStatisticsData `json:"data"`
}

func (c *Client) DuowanStatistics(ctx context.Context, server string, token string) (*[]DuowanStatisticsServer, error) {
	params := &struct {
		Server string `json:"server"`
		Token  string `json:"token"`
	}{
		Server: server,
		Token:  token,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("DuowanStatistics: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/duowan/statistics", body)
	if err != nil {
		slog.Error("DuowanStatistics: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("DuowanStatistics: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]DuowanStatisticsServer)

	if resp.Msg != "success" {
		slog.Error("DuowanStatistics: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("DuowanStatistics: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
