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

func (c *Client) RanchRecords(ctx context.Context, server string, token string) (*RanchResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Token  string `json:"token"`
	}{
		Server: server,
		Token:  token,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("RanchRecords: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/ranch/records", body)
	if err != nil {
		slog.Error("HorseRanch: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("RanchRecords: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(RanchResponse)

	if resp.Msg != "success" {
		slog.Error("RanchRecords: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("RanchRecords: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
