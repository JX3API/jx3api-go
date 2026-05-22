package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type DemonPriceResponse struct {
	Zone       string `json:"zone"`
	Server     string `json:"server"`
	Tieba      string `json:"tieba"`
	Wanbaolou  string `json:"wanbaolou"`
	Dd373      string `json:"dd373"`
	Date       string `json:"date"`
}

func (c *Client) TradeDemon(ctx context.Context, server string, limit int, token string) (*[]DemonPriceResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Limit  int    `json:"limit"`
		Token  string `json:"token"`
	}{
		Server: server,
		Limit:  limit,
		Token:  token,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("TradeDemon: request body marshal error: " + err.Error())
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
		slog.Info("TradeDemon: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]DemonPriceResponse)

	if resp.Msg != "success" {
		slog.Error("TradeDemon: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("TradeDemon: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
