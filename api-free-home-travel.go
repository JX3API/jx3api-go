package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type HomeTravelResponse struct {
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

func (c *Client) HomeTravel(ctx context.Context, name string) (*[]HomeTravelResponse, error) {
	params := &struct {
		Name string `json:"name"`
	}{
		Name: name,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("HomeTravel: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/home/travel", body)
	if err != nil {
		slog.Error("HomeTravel: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("HomeTravel: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]HomeTravelResponse)

	if resp.Msg != "success" {
		slog.Error("HomeTravel: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("HomeTravel: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
