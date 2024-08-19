package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type TiebaItemRecordResponse struct {
	ID      int    `json:"id"`
	Zone    string `json:"zone"`
	Server  string `json:"server"`
	Name    string `json:"name"`
	URL     int    `json:"url"`
	Context string `json:"context"`
	Reply   int    `json:"reply"`
	Token   string `json:"token"`
	Floor   int    `json:"floor"`
	Time    int    `json:"time"`
}

func (c *Client) TiebaItemRecord(ctx context.Context, name string, server string, limit int) (*[]TiebaItemRecordResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Name   string `json:"name"`
		Limit  int    `json:"limit"`
	}{
		Name:  name,
		Limit: limit,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("TiebaItemRecord: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/tieba/item/record", body)
	if err != nil {
		slog.Error("TiebaItemRecord: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("TiebaItemRecord: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]TiebaItemRecordResponse)

	if resp.Msg != "success" {
		slog.Error("TiebaItemRecord: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("TiebaItemRecord: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
