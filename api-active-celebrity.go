package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type ActivateCelebrityResponse struct {
	MapName string `json:"map_name"`
	Event   string `json:"event"`
	Site    string `json:"site"`
	Desc    string `json:"desc"`
	Icon    string `json:"icon"`
	Time    string `json:"time"`
}

func (c *Client) ActivateCelebrity(ctx context.Context, season int) (*[]ActivateCelebrityResponse, error) {
	params := struct {
		Season int `json:"season"`
	}{
		Season: season,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("ActivateCelebrity: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/active/celebrity", body)
	if err != nil {
		slog.Error("ActivateCelebrity: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("ActivateCelebrity: response body unmarshal error: ", err)
		return nil, err
	}

	data := new([]ActivateCelebrityResponse)

	if resp.Msg != "success" {
		slog.Error("ActivateCelebrity: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("ActivateCelebrity: data unmarshal error: ", err)
		return nil, err
	}

	return data, nil
}
