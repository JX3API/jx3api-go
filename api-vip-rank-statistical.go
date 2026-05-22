package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type RankStatisticalData struct {
	MaxCount    int    `json:"max_count"`
	NowCount    int    `json:"now_count"`
	TongName    string `json:"tong_name"`
	CastleName  string `json:"castle_name"`
	MasterName  string `json:"master_name"`
	TotalScore  int    `json:"total_score"`
}

type RankStatisticalResponse struct {
	Id       int                  `json:"id"`
	Zone     string               `json:"zone"`
	Server   string               `json:"server"`
	Name     string               `json:"name"`
	Data     []RankStatisticalData `json:"data"`
	Time     int64                `json:"time"`
}

func (c *Client) RankStatistical(ctx context.Context, server string, name string, token string) (*RankStatisticalResponse, error) {
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
		slog.Error("RankStatistical: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/rank/statistics", body)
	if err != nil {
		slog.Error("RankStatistical: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("RankStatistical: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(RankStatisticalResponse)

	if resp.Msg != "success" {
		slog.Error("RankStatistical: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("RankStatistical: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
