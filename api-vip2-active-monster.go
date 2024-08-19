package jx3api

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type MonsterResponse struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
	Data  []struct {
		Level int      `json:"level"`
		Name  string   `json:"name"`
		Skill []string `json:"skill"`
		Data  struct {
			Name string   `json:"name"`
			List []string `json:"list"`
			Desc string   `json:"desc"`
		} `json:"data"`
	} `json:"data"`
}

func (c *Client) ActiveMonster(ctx context.Context) (*MonsterResponse, error) {
	raw, err := c.request(ctx, "/data/active/monster", nil)
	if err != nil {
		slog.Error("ActiveMonster: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("ActiveMonster: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(MonsterResponse)

	if resp.Msg != "success" {
		slog.Error("ActiveMonster: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("ActiveMonster: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
