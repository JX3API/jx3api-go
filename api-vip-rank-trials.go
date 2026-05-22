package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type TrialsData struct {
	MaxLevel   int    `json:"max_level"`
	RoleName   string `json:"role_name"`
	EquipScore int    `json:"equip_score"`
	TotalScore int    `json:"total_score"`
}

type TrialsResponse struct {
	ID       int          `json:"id"`
	Zone     string       `json:"zone"`
	Server   string       `json:"server"`
	Name     string       `json:"name"`
	Data     []TrialsData `json:"data"`
	Time     int64        `json:"time"`
}

func (c *Client) RankTrials(ctx context.Context, server string, name string, token string) (*TrialsResponse, error) {
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
		slog.Error("RankTrials: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/rank/trials", body)
	if err != nil {
		slog.Error("RankTrials: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("RankTrials: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(TrialsResponse)

	if resp.Msg != "success" {
		slog.Error("RankTrials: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("RankTrials: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
