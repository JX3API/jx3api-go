package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type CardRandomResponse struct {
	ZoneName   string `json:"zoneName"`
	ServerName string `json:"serverName"`
	RoleName   string `json:"roleName"`
	ShowHash   string `json:"showHash"`
	ShowIndex  int    `json:"showIndex"`
	ShowAvatar string `json:"showAvatar"`
}

func (c *Client) CardRandom(ctx context.Context, server string, body string, force string, token string) (*CardRandomResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Body   string `json:"body"`
		Force  string `json:"force"`
		Token  string `json:"token"`
	}{
		Server: server,
		Body:   body,
		Force:  force,
		Token:  token,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("CardRandom: request body marshal error: " + err.Error())
		return nil, err
	}

	bodyReader := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/card/random", bodyReader)
	if err != nil {
		slog.Error("CardRandom: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("CardRandom: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(CardRandomResponse)

	if resp.Msg != "success" {
		slog.Error("CardRandom: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("CardRandom: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
