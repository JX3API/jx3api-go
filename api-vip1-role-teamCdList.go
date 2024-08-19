package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type TeamCdListResponse struct {
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
	Data         []struct {
		MapIcon      string `json:"mapIcon"`
		MapId        string `json:"mapId"`
		MapName      string `json:"mapName"`
		MapType      string `json:"mapType"`
		BossCount    int    `json:"bossCount"`
		BossFinished int    `json:"bossFinished"`
		BossProgress []struct {
			Finished   bool   `json:"finished"`
			Icon       string `json:"icon"`
			Index      string `json:"index"`
			Name       string `json:"name"`
			ProgressId string `json:"progressId"`
		} `json:"bossProgress"`
	} `json:"data"`
}

func (c *Client) RoleTeamCdList(ctx context.Context, server string, name string) (*TeamCdListResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Name   string `json:"name"`
		Ticket string `json:"ticket"`
	}{
		Server: server,
		Name:   name,
		Ticket: c.Opts.Ticket,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("TeamCdList: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/role/teamCdList", body)
	if err != nil {
		slog.Error("TeamCdList: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("TeamCdList: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(TeamCdListResponse)

	if resp.Msg != "success" {
		slog.Error("TeamCdList: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("TeamCdList: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
