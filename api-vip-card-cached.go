package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type CardCachedResponse struct {
	ZoneName   string `json:"zoneName"`
	ServerName string `json:"serverName"`
	RoleName   string `json:"roleName"`
	ShowHash   string `json:"showHash"`
	ShowAvatar string `json:"showAvatar"`
}

func (c *Client) CardCached(ctx context.Context, server string, name string, token string) (*CardCachedResponse, error) {
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
		slog.Error("CardCached: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/card/cached", body)
	if err != nil {
		slog.Error("CardCached: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("CardCached: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(CardCachedResponse)

	if resp.Msg != "success" {
		slog.Error("CardCached: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("CardCached: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
