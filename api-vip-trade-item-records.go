package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type TradeItemRecordItem struct {
	ID     string `json:"id"`
	Index  int    `json:"index"`
	Zone   string `json:"zone"`
	Server string `json:"server"`
	Value  int    `json:"value"`
	Sale   int    `json:"sale"`
	Token  string `json:"token"`
	Date   string `json:"date"`
	Source int    `json:"source"`
	Status int    `json:"status"`
}

type TradeItemRecordsData struct {
	Class   string                `json:"class"`
	Subclass string               `json:"subclass"`
	Name    string               `json:"name"`
	Alias   string               `json:"alias"`
	Value   string               `json:"value"`
	Desc    string               `json:"desc"`
	Date    string               `json:"date"`
	View    string               `json:"view"`
	List    [][]TradeItemRecordItem `json:"list"`
}

type TradeItemRecordsResponse struct {
	Data TradeItemRecordsData `json:"data"`
	Time int64                `json:"time"`
}

func (c *Client) TradeItemRecords(ctx context.Context, server string, name string, token string) (*TradeItemRecordsResponse, error) {
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
		slog.Error("TradeItemRecords: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/trade/item/records", body)
	if err != nil {
		slog.Error("TradeItemRecords: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("TradeItemRecords: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(TradeItemRecordsResponse)

	if resp.Msg != "success" {
		slog.Error("TradeItemRecords: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("TradeItemRecords: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
