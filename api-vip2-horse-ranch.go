package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type RanchResponse struct {
	Zone   string              `json:"zone"`
	Server string              `json:"server"`
	Data   map[string][]string `json:"data"`
	Note   string              `json:"note"`
}

func (c *Client) HorseRanch(ctx context.Context, server string) (*RanchResponse, error) {
	params := &struct {
		Server string `json:"server"`
	}{
		Server: server,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("HorseRanch: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/horse/ranch", body)
	if err != nil {
		slog.Error("HorseRanch: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("HorseRanch: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(RanchResponse)

	if resp.Msg != "success" {
		slog.Error("HorseRanch: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("HorseRanch: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
