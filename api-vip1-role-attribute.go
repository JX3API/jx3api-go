package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
)

func (c *Client) RoleAttribute(ctx context.Context, server string, name string) ([]byte, error) {
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
		slog.Error("RoleAttribute: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/role/detailed", body)
	if err != nil {
		slog.Error("RoleAttribute: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("RoleAttribute: response body unmarshal error: " + err.Error())
		return nil, err
	}

	return raw, nil

}
