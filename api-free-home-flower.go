package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type FlowerResponse struct {
	Name  string   `json:"name"`
	Color string   `json:"color"`
	Price float64  `json:"price"`
	Line  []string `json:"line"`
}

func (c *Client) HomeFlower(ctx context.Context, server, name, mapName string) (*map[string][]FlowerResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Name   string `json:"name"`
		Map    string `json:"map"`
	}{
		Server: server,
		Name:   name,
		Map:    mapName,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("Flower: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/home/flower", body)
	if err != nil {
		slog.Error("Flower: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("Flower: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(map[string][]FlowerResponse)

	if resp.Msg != "success" {
		slog.Error("Flower: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("Flower: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
