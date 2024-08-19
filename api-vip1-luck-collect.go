package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type LuckCollectResponse struct {
	Server string `json:"server"`
	Event  string `json:"event"`
	Count  int    `json:"count"`
	Data   struct {
		Name string `json:"name"`
		Time int64  `json:"time"`
	} `json:"data"`
}

func (c *Client) LuckCollect(ctx context.Context, server string, num int) (*[]LuckCollectResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Num    int    `json:"num"`
	}{
		Server: server,
		Num:    num,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("LuckCollect: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/luck/collect", body)
	if err != nil {
		slog.Error("LuckCollect: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("LuckCollect: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]LuckCollectResponse)

	if resp.Msg != "success" {
		slog.Error("LuckCollect: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("LuckCollect: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
