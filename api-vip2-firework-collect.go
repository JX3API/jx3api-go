package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type FireworkCollectResponse struct {
	Server    string `json:"server"`
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Name      string `json:"name"`
	Count     int    `json:"count"`
	Time      int64  `json:"time"`
}

func (c *Client) FireworkCollect(ctx context.Context, server string, num int) (*[]FireworkCollectResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Num    int    `json:"num"`
	}{
		Server: server,
		Num:    num,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("FireworkCollect: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/firework/collect", body)
	if err != nil {
		slog.Error("FireworkCollect: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("FireworkCollect: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]FireworkCollectResponse)

	if resp.Msg != "success" {
		slog.Error("FireworkCollect: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("FireworkCollect: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
