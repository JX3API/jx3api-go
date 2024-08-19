package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type AdventureResponse struct {
	Zone   string `json:"zone"`
	Server string `json:"server"`
	Name   string `json:"name"`
	Event  string `json:"event"`
	Level  int    `json:"level"`
	Status int    `json:"status"`
	Time   int64  `json:"time"`
}

func (c *Client) LuckAdventure(ctx context.Context, server string, name string) (*[]AdventureResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Name   string `json:"name"`
		Ticket string `json:"ticket"`
	}{
		Server: server,
		Name:   name,
		Ticket: c.Opts.Ticket,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("RoleAdventure: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/luck/adventure", body)
	if err != nil {
		slog.Error("RoleAdventure: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("RoleAdventure: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]AdventureResponse)

	if resp.Msg != "success" {
		slog.Error("RoleAdventure: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("RoleAdventure: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
