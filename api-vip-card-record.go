package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type CardRecordResponse struct {
	ZoneName   string `json:"zoneName"`
	ServerName string `json:"serverName"`
	RoleName   string `json:"roleName"`
	ShowHash   string `json:"showHash"`
	ShowIndex  int    `json:"showIndex"`
	ShowAvatar string `json:"showAvatar"`
	CacheTime  int64  `json:"cacheTime"`
}

func (c *Client) CardRecord(ctx context.Context, server string, name string, token string) (*CardRecordResponse, error) {
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
		slog.Error("CardRecord: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/card/record", body)
	if err != nil {
		slog.Error("CardRecord: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("CardRecord: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(CardRecordResponse)

	if resp.Msg != "success" {
		slog.Error("CardRecord: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("CardRecord: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
