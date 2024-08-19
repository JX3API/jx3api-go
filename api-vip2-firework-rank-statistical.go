package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type FireworkRankResponse struct {
	Server    string `json:"server"`
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Name      string `json:"name"`
	Count     int    `json:"count"`
	Time      int64  `json:"time"`
}

func (c *Client) FireworkRankStatistical(ctx context.Context, server string, column string, thisTime int64, thatTime int64) (*[]FireworkRankResponse, error) {
	params := &struct {
		Server   string `json:"server"`
		Column   string `json:"column"`
		ThisTime int64  `json:"this_time"`
		ThatTime int64  `json:"that_time"`
	}{
		Server:   server,
		Column:   column,
		ThisTime: thisTime,
		ThatTime: thatTime,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("FireworkRankStatistical: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/firework/rank/statistical", body)
	if err != nil {
		slog.Error("FireworkRankStatistical: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("FireworkRankStatistical: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]FireworkRankResponse)

	if resp.Msg != "success" {
		slog.Error("FireworkRankStatistical: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("FireworkRankStatistical: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
