package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type ValuablesStatisticalResponse struct {
	ID       int    `json:"id"`
	Zone     string `json:"zone"`
	Server   string `json:"server"`
	Name     string `json:"name"`
	RoleName string `json:"role_name"`
	MapName  string `json:"map_name"`
	Time     int64  `json:"time"`
}

func (c *Client) ValuablesStatistical(ctx context.Context, server string, name string, limit int) (*[]ValuablesStatisticalResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Name   string `json:"name"`
		Limit  int    `json:"limit"`
	}{
		Server: server,
		Name:   name,
		Limit:  limit,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("ValuablesStatistical: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/valuables/statistical", body)
	if err != nil {
		slog.Error("ValuablesStatistical: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("ValuablesStatistical: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]ValuablesStatisticalResponse)

	if resp.Msg != "success" {
		slog.Error("ValuablesStatistical: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("ValuablesStatistical: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
