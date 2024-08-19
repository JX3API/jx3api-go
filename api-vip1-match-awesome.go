package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type AwesomeListResponse struct {
	ZoneName   string `json:"zoneName"`
	ServerName string `json:"serverName"`
	RoleName   string `json:"roleName"`
	ForceName  string `json:"forceName"`
	AvatarUrl  string `json:"avatarUrl"`
	RankNum    string `json:"rankNum"`
	Score      string `json:"score"`
	UpNum      string `json:"upNum"`
	WinRate    string `json:"winRate"`
}

func (c *Client) MatchAwesome(ctx context.Context, mode int, limit int) (*[]AwesomeListResponse, error) {
	params := &struct {
		Mode   int    `json:"mode"`
		Limit  int    `json:"limit"`
		Ticket string `json:"ticket"`
	}{
		Mode:   mode,
		Limit:  limit,
		Ticket: c.Opts.Ticket,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("MatchAwesome: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/match/awesome", body)
	if err != nil {
		slog.Error("MatchAwesome: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("MatchAwesome: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]AwesomeListResponse)

	if resp.Msg != "success" {
		slog.Error("MatchAwesome: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("MatchAwesome: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
