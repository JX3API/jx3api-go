package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
)

type ArenaRecentResponse struct {
	BodyId    int    `json:"bodyId"`
	BodyName  string `json:"bodyName"`
	CampId    int    `json:"campId"`
	CampName  string `json:"campName"`
	ForceId   int    `json:"forceId"`
	ForceName string `json:"forceName"`
	GlobalId  string `json:"globalId"`
	History   []struct {
		AvgGrade  int    `json:"avgGrade"`
		EndTime   int64  `json:"endTime"`
		Kungfu    string `json:"kungfu"`
		Mmr       int    `json:"mmr"`
		Mvp       bool   `json:"mvp"`
		PvpType   int    `json:"pvpType"`
		StartTime int64  `json:"startTime"`
		TotalMmr  int    `json:"totalMmr"`
		Won       bool   `json:"won"`
	} `json:"history"`
	Performance struct {
		V2v2 ArenaPerformance `json:"2v2"`
		V3v3 ArenaPerformance `json:"3v3"`
		V5v5 ArenaPerformance `json:"5v5"`
	} `json:"performance"`
	RoleId     string `json:"roleId"`
	RoleName   string `json:"roleName"`
	ServerName string `json:"serverName"`
	TongId     int    `json:"tongId"`
	TongName   string `json:"tongName"`
	ZoneName   string `json:"zoneName"`
}

type ArenaPerformance struct {
	Mmr        int    `json:"mmr"`
	Grade      int    `json:"grade"`
	Ranking    string `json:"ranking"`
	WinCount   int    `json:"winCount"`
	TotalCount int    `json:"totalCount"`
	MvpCount   int    `json:"mvpCount"`
	PvpType    string `json:"pvpType"`
	WinRate    int    `json:"winRate"`
}

func (p *ArenaPerformance) UnmarshalJSON(raw []byte) error {
	if bytes.Equal(bytes.TrimSpace(raw), []byte("[]")) {
		*p = ArenaPerformance{}
		return nil
	}
	type plain ArenaPerformance
	return json.Unmarshal(raw, (*plain)(p))
}

type ArenaRecentPerformance2v2 = ArenaPerformance
type ArenaRecentPerformance3v3 = ArenaPerformance
type ArenaRecentPerformance5v5 = ArenaPerformance

// ArenaRecent 查询角色名剑大会表现及近期战绩。
func (c *Client) ArenaRecent(ctx context.Context, server string, name string, mode int, ticket string, token string) (*ArenaRecentResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Name   string `json:"name"`
		Mode   int    `json:"mode"`
		Ticket string `json:"ticket,omitempty"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Name:   name,
		Mode:   mode,
		Ticket: ticket,
		Token:  token,
	}
	return requestData[ArenaRecentResponse](ctx, c, "/arena/recent", params)
}
