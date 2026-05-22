package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type MonsterResponse struct {
	Week string `json:"week"`
	Start int64 `json:"start"`
	End   int64 `json:"end"`
	Boss  string `json:"boss"`
	List  []struct {
		Index int      `json:"index"`
		Name  string   `json:"name"`
		Skill []string `json:"skill"`
		Data  struct {
			Name string   `json:"name"`
			List []string `json:"list"`
			Desc string   `json:"desc"`
		} `json:"data"`
	} `json:"list"`
}

func (c *Client) ActiveMonster(ctx context.Context, token string) (*MonsterResponse, error) {
	params := &struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("ActiveMonster: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/active/monster", body)
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
