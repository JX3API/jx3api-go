package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type DemonPriceResponse struct {
	Id         int    `json:"id"`
	Zone       string `json:"zone"`
	Server     string `json:"server"`
	Tieba      string `json:"tieba"`
	Wanbaolou  string `json:"wanbaolou"`
	Gold_dd373 string `json:"dd373"`
	Gold_uu898 string `json:"uu898"`
	Gold_5173  string `json:"5173"`
	Gold_7881  string `json:"7881"`
	Time       int64  `json:"time"`
	Date       string `json:"date"`
}

func (c *Client) DemonPrice(ctx context.Context, server string, limit int) (*[]DemonPriceResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Limit  int    `json:"limit"`
	}{
		Server: server,
		Limit:  limit,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("DemonPrice: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/trade/demon", body)
	if err != nil {
		slog.Error("DemonPrice: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("DemonPrice: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]DemonPriceResponse)

	if resp.Msg != "success" {
		slog.Error("DemonPrice: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("DemonPrice: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
