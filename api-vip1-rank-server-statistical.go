package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type ServerRankResponse struct {
	Id        int    `json:"id"`
	Class     string `json:"class"`
	Zone      string `json:"zone"`
	Server    string `json:"server"`
	Name      string `json:"name"`
	School    string `json:"school"`
	Index     int    `json:"index"`
	Level     int    `json:"level"`
	CampName  string `json:"camp_name"`
	TongName  string `json:"tong_name"`
	Score     int    `json:"score"`
	Datetime  string `json:"datetime"`
	GuildName string `json:"guild_name"`
	GuildId   int    `json:"guild_id"`
	WinCount  int    `json:"win_count"`
	LoseCount int    `json:"lose_count"`
}

func (c *Client) ServerRank(ctx context.Context, table string, name string) (*[]ServerRankResponse, error) {
	params := &struct {
		Table string `json:"table"`
		Name  string `json:"name"`
	}{
		Table: table,
		Name:  name,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("ServerRank: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/rank/server/statistical", body)
	if err != nil {
		slog.Error("ServerRank: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("ServerRank: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]ServerRankResponse)

	if resp.Msg != "success" {
		slog.Error("ServerRank: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("ServerRank: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
