package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type TradeRecordItem struct {
	ID       string `json:"id"`
	Index    int    `json:"index"`
	Zone     string `json:"zone"`
	Server   string `json:"server"`
	Value    int    `json:"value"`
	Sale     int    `json:"sale"`
	Token    string `json:"token"`
	Date     string `json:"date"`
	Source   int    `json:"source"`
	Status   int    `json:"status"`
}

type TradeRecordsData struct {
	Class   string             `json:"class"`
	Subclass string            `json:"subclass"`
	Name    string             `json:"name"`
	Alias   string             `json:"alias"`
	Value   string             `json:"value"`
	Desc    string             `json:"desc"`
	Date    string             `json:"date"`
	View    string             `json:"view"`
	List    [][]TradeRecordItem `json:"list"`
}

type TradeRecordsResponse struct {
	Data TradeRecordsData `json:"data"`
	Time int64            `json:"time"`
}

func (c *Client) TradeRecords(ctx context.Context, server string, name string, token string) (*TradeRecordsResponse, error) {
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
		slog.Error("TradeRecords: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/trade/records", body)
	if err != nil {
		slog.Error("TradeRecords: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("TradeRecords: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(TradeRecordsResponse)

	if resp.Msg != "success" {
		slog.Error("TradeRecords: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("TradeRecords: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
