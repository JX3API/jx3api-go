package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type BattleRecord struct {
	ZoneName         string `json:"zoneName"`
	ServerName       string `json:"serverName"`
	DeclaringTongName string `json:"declaringTongName"`
	AcceptingTongName string `json:"acceptingTongName"`
	StartTime        int64  `json:"startTime"`
	MatchDuration    int    `json:"matchDuration"`
	EndTime          int64  `json:"endTime"`
}

func (c *Client) BattleRecords(ctx context.Context, server string, token string) (*[]BattleRecord, error) {
	params := &struct {
		Server string `json:"server"`
		Token  string `json:"token"`
	}{
		Server: server,
		Token:  token,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("BattleRecords: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/battle/records", body)
	if err != nil {
		slog.Error("BattleRecords: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("BattleRecords: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]BattleRecord)

	if resp.Msg != "success" {
		slog.Error("BattleRecords: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("BattleRecords: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
