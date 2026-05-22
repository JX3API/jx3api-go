package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type CardRecordItem struct {
	ZoneName   string `json:"zoneName"`
	ServerName string `json:"serverName"`
	RoleName   string `json:"roleName"`
	ShowHash   string `json:"showHash"`
	ShowIndex  int    `json:"showIndex"`
	ShowActive bool   `json:"showActive"`
	ShowAvatar string `json:"showAvatar"`
	SaveTime   int64  `json:"saveTime"`
}

type CardRecordsResponse struct {
	Data []CardRecordItem `json:"data"`
}

func (c *Client) CardRecords(ctx context.Context, server string, name string, token string) (*[]CardRecordItem, error) {
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
		slog.Error("CardRecords: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/card/records", body)
	if err != nil {
		slog.Error("CardRecords: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("CardRecords: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]CardRecordItem)

	if resp.Msg != "success" {
		slog.Error("CardRecords: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("CardRecords: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
