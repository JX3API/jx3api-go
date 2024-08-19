package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type HomeFurnitureResponse struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Type         int    `json:"type"`
	Color        int    `json:"color"`
	Source       string `json:"source"`
	Architecture int    `json:"architecture"`
	Limit        int    `json:"limit"`
	Quality      int    `json:"quality"`
	View         int    `json:"view"`
	Practical    int    `json:"practical"`
	Hard         int    `json:"hard"`
	Geomantic    int    `json:"geomantic"`
	Interesting  int    `json:"interesting"`
	Produce      string `json:"produce"`
	Image        string `json:"image"`
	Tip          string `json:"tip"`
}

func (c *Client) HomeFurniture(ctx context.Context, name string) (*HomeFurnitureResponse, error) {
	params := &struct {
		Name string `json:"name"`
	}{
		Name: name,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("HomeFurniture: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/home/furniture", body)
	if err != nil {
		slog.Error("HomeFurniture: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("HomeFurniture: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(HomeFurnitureResponse)

	if resp.Msg != "success" {
		slog.Error("HomeFurniture: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("HomeFurniture: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
