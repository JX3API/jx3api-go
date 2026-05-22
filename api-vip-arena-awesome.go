package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type ArenaAwesomeResponse struct {
	ZoneName  string `json:"zoneName"`
	ServerName string `json:"serverName"`
	RoleName  string `json:"roleName"`
	ForceName string `json:"forceName"`
	AvatarUrl string `json:"avatarUrl"`
	RankNum   string `json:"rankNum"`
	Score     string `json:"score"`
	UpNum     string `json:"upNum"`
	WinRate   string `json:"winRate"`
}

func (c *Client) ArenaAwesome(ctx context.Context, mode int, limit int, ticket string, token string) (*[]ArenaAwesomeResponse, error) {
	params := &struct {
		Mode   int    `json:"mode"`
		Limit  int    `json:"limit"`
		Ticket string `json:"ticket"`
		Token  string `json:"token"`
	}{
		Mode:   mode,
		Limit:  limit,
		Ticket: ticket,
		Token:  token,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("ArenaAwesome: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/arena/awesome", body)
	if err != nil {
		slog.Error("ArenaAwesome: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("ArenaAwesome: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]ArenaAwesomeResponse)

	if resp.Msg != "success" {
		slog.Error("ArenaAwesome: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("ArenaAwesome: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
