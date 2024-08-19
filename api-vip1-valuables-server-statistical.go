package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type ValuablesServerStatisticalResponse struct {
	ID       int    `json:"id"`
	Zone     string `json:"zone"`
	Server   string `json:"server"`
	Name     string `json:"name"`
	RoleName string `json:"role_name"`
	MapName  string `json:"map_name"`
	Time     int64  `json:"time"`
}

func (c *Client) ValuablesServerStatistical(ctx context.Context, name string, limit int) (*[]ValuablesServerStatisticalResponse, error) {
	params := &struct {
		Name  string `json:"name"`
		Limit int    `json:"limit"`
	}{
		Name:  name,
		Limit: limit,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("ValuablesServerStatistical: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/valuables/server/statistical", body)
	if err != nil {
		slog.Error("ValuablesServerStatistical: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("ValuablesServerStatistical: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]ValuablesServerStatisticalResponse)

	if resp.Msg != "success" {
		slog.Error("ValuablesServerStatistical: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("ValuablesServerStatistical: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
