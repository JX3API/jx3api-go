package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type MatchRecentResponse struct {
	ZoneName     string `json:"zoneName"`
	ServerName   string `json:"serverName"`
	RoleName     string `json:"roleName"`
	RoleId       string `json:"roleId"`
	GlobalRoleId string `json:"globalRoleId"`
	ForceName    string `json:"forceName"`
	ForceId      string `json:"forceId"`
	BodyName     string `json:"bodyName"`
	BodyId       string `json:"bodyId"`
	TongName     string `json:"tongName"`
	TongId       string `json:"tongId"`
	CampName     string `json:"campName"`
	CampId       string `json:"campId"`
	PersonName   string `json:"personName"`
	PersonId     string `json:"personId"`
	PersonAvatar string `json:"personAvatar"`
	Performance  struct {
		Type_2v2 struct {
			Mmr        int    `json:"mmr"`
			Grade      int    `json:"grade"`
			Ranking    string `json:"ranking"`
			WinCount   int    `json:"winCount"`
			TotalCount int    `json:"totalCount"`
			MvpCount   int    `json:"mvpCount"`
			PvpType    string `json:"pvpType"`
			WinRate    int    `json:"winRate"`
		} `json:"2v2,omitempty"`
		Type_3v3 struct {
			Mmr        int    `json:"mmr"`
			Grade      int    `json:"grade"`
			Ranking    string `json:"ranking"`
			WinCount   int    `json:"winCount"`
			TotalCount int    `json:"totalCount"`
			MvpCount   int    `json:"mvpCount"`
			PvpType    string `json:"pvpType"`
			WinRate    int    `json:"winRate"`
		} `json:"3v3,omitempty"`
		Type_5v5 struct {
			Mmr        int    `json:"mmr"`
			Grade      int    `json:"grade"`
			Ranking    string `json:"ranking"`
			WinCount   int    `json:"winCount"`
			TotalCount int    `json:"totalCount"`
			MvpCount   int    `json:"mvpCount"`
			PvpType    string `json:"pvpType"`
			WinRate    int    `json:"winRate"`
		} `json:"5v5,omitempty"`
	} `json:"performance"`
	History []struct {
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
	} `json:"history"`
	Trend []struct {
		MatchDate int64   `json:"matchDate"`
		Mmr       int     `json:"mmr"`
		WinRate   float64 `json:"winRate"`
	} `json:"trend"`
}

func (c *Client) MatchRecent(ctx context.Context, server string, name string, mode int) (*MatchRecentResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Name   string `json:"name"`
		Mode   int    `json:"mode"`
		Ticket string `json:"ticket"`
	}{
		Server: server,
		Name:   name,
		Mode:   mode,
		Ticket: c.Opts.Ticket,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("MatchRecent: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/match/recent", body)
	if err != nil {
		slog.Error("MatchRecent: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("MatchRecent: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(MatchRecentResponse)

	if resp.Msg != "success" {
		slog.Error("MatchRecent: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("MatchRecent: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
