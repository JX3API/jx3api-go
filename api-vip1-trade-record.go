package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type TradeRecordResponse struct {
	ID       int    `json:"id"`
	Class    string `json:"class"`
	Subclass string `json:"subclass"`
	Name     string `json:"name"`
	Alias    string `json:"alias"`
	Subalias string `json:"subalias"`
	Raw      string `json:"raw"`
	Level    int    `json:"level"`
	Desc     string `json:"desc"`
	View     string `json:"view"`
	Date     string `json:"date"`
	Data     [][]struct {
		ID       string `json:"id"`
		Index    int    `json:"index"`
		Zone     string `json:"zone"`
		Server   string `json:"server"`
		Value    int    `json:"value"`
		Sales    int    `json:"sales"`
		Token    string `json:"token"`
		Source   int    `json:"source"`
		Date     string `json:"date"`
		Status   int    `json:"status"`
		Datetime string `json:"datetime"`
	} `json:"data"`
}

func (c *Client) TradeRecord(ctx context.Context, name string, server ...string) (*TradeRecordResponse, error) {
	params := &struct {
		Server string `json:"server,omitempty"`
		Name   string `json:"name"`
	}{
		Name: name,
	}

	if len(server) > 0 {
		params.Server = server[0]
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("TradeRecord: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/trade/record", body)
	if err != nil {
		slog.Error("TradeRecord: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("TradeRecord: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(TradeRecordResponse)

	if resp.Msg != "success" {
		slog.Error("TradeRecord: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("TradeRecord: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
