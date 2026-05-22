package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type EventCollectData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Time int64  `json:"time"`
}

type EventCollectItem struct {
	Zone  string           `json:"zone"`
	Server string          `json:"server"`
	Event string           `json:"event"`
	Count int              `json:"count"`
	Data  EventCollectData `json:"data"`
}

type EventCollectResponse struct {
	Items []EventCollectItem `json:"data"`
}

func (c *Client) EventCollect(ctx context.Context, server string, num int, token string) (*[]EventCollectItem, error) {
	params := &struct {
		Server string `json:"server"`
		Num    int    `json:"num"`
		Token  string `json:"token"`
	}{
		Server: server,
		Num:    num,
		Token:  token,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("EventCollect: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/event/collect", body)
	if err != nil {
		slog.Error("EventCollect: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("EventCollect: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]EventCollectItem)

	if resp.Msg != "success" {
		slog.Error("EventCollect: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("EventCollect: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
