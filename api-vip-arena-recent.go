package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type ArenaRecentPerformance3v3 struct {
	Mmr       int    `json:"mmr"`
	Grade     int    `json:"grade"`
	Ranking   string `json:"ranking"`
	WinCount  int    `json:"winCount"`
	TotalCount int   `json:"totalCount"`
	MvpCount  int    `json:"mvpCount"`
	PvpType   string `json:"pvpType"`
	WinRate   int    `json:"winRate"`
}

type ArenaRecentPerformance2v2 struct {
	Mmr       int    `json:"mmr"`
	Grade     int    `json:"grade"`
	Ranking   string `json:"ranking"`
	WinCount  int    `json:"winCount"`
	TotalCount int   `json:"totalCount"`
	MvpCount  int    `json:"mvpCount"`
	PvpType   string `json:"pvpType"`
	WinRate   int    `json:"winRate"`
}

type ArenaRecentPerformance5v5 []interface{}

type ArenaRecentHistory struct {
	Zone      string `json:"zone"`
	Server    string `json:"server"`
	AvgGrade  int    `json:"avgGrade"`
	TotalMmr  int    `json:"totalMmr"`
	Mmr       int    `json:"mmr"`
	Kungfu    string `json:"kungfu"`
	PvpType   int    `json:"pvpType"`
	Won       bool   `json:"won"`
	Mvp       bool   `json:"mvp"`
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
}

type ArenaRecentResponse struct {
	ZoneName   string                      `json:"zoneName"`
	ServerName string                     `json:"serverName"`
	RoleName   string                     `json:"roleName"`
	RoleId     string                     `json:"roleId"`
	GlobalId   string                     `json:"globalId"`
	ForceName  string                     `json:"forceName"`
	ForceId    int                        `json:"forceId"`
	BodyName   string                     `json:"bodyName"`
	BodyId     int                        `json:"bodyId"`
	TongName   string                     `json:"tongName"`
	TongId     int                        `json:"tongId"`
	CampName   string                     `json:"campName"`
	CampId     string                     `json:"campId"`
	Performance struct {
		V5v5 ArenaRecentPerformance5v5              `json:"5v5"`
		V3v3 ArenaRecentPerformance3v3              `json:"3v3"`
		V2v2 ArenaRecentPerformance2v2              `json:"2v2"`
	} `json:"performance"`
	History []ArenaRecentHistory `json:"history"`
}

func (c *Client) ArenaRecent(ctx context.Context, server string, name string, mode int, ticket string, token string) (*ArenaRecentResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Name   string `json:"name"`
		Mode   int    `json:"mode"`
		Ticket string `json:"ticket"`
		Token  string `json:"token"`
	}{
		Server: server,
		Name:   name,
		Mode:   mode,
		Ticket: ticket,
		Token:  token,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("ArenaRecent: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/arena/recent", body)
	if err != nil {
		slog.Error("ArenaRecent: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("ArenaRecent: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(ArenaRecentResponse)

	if resp.Msg != "success" {
		slog.Error("ArenaRecent: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("ArenaRecent: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
