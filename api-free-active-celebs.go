package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type ActiveCelebsResponse struct {
	MapName string `json:"map_name"`
	Event   string `json:"event"`
	Site    string `json:"site"`
	Desc    string `json:"desc"`
	Icon    string `json:"icon"`
	Time    string `json:"time"`
}

func (c *Client) ActiveCelebs(ctx context.Context, name string) (*[]ActiveCelebsResponse, error) {
	params := &struct {
		Name string `json:"name"`
	}{
		Name: name,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("ActiveCelebs: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/active/celebs", body)
	if err != nil {
		slog.Error("ActiveCelebs: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("ActiveCelebs: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]ActiveCelebsResponse)

	if resp.Msg != "success" {
		slog.Error("ActiveCelebs: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("ActiveCelebs: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
