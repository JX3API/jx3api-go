package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type TradeItemSearchResponse struct {
	Class    string `json:"class"`
	Subclass string `json:"subclass"`
	Name     string `json:"name"`
	Alias    string `json:"alias"`
	Wblalias string `json:"wblalias"`
	Value    string `json:"value"`
	Desc     string `json:"desc"`
	Date     string `json:"date"`
	View     string `json:"view"`
}

func (c *Client) TradeItemSearch(ctx context.Context, name string, token string) (*[]TradeItemSearchResponse, error) {
	params := &struct {
		Name  string `json:"name"`
		Token string `json:"token"`
	}{
		Name:  name,
		Token: token,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("TradeItemSearch: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/trade/item/search", body)
	if err != nil {
		slog.Error("TradeItemSearch: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("TradeItemSearch: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]TradeItemSearchResponse)

	if resp.Msg != "success" {
		slog.Error("TradeItemSearch: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("TradeItemSearch: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
