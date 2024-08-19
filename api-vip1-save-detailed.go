package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type SaveDetailedResponse struct {
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
}

func (c *Client) SaveDetailed(ctx context.Context, server string, roleId string) (*SaveDetailedResponse, error) {
	params := &struct {
		Server string `json:"server"`
		RoleId string `json:"roleid"`
		Ticket string `json:"ticket"`
	}{
		Server: server,
		RoleId: roleId,
		Ticket: c.Opts.Ticket,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("SaveDetailed: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/save/detailed", body)
	if err != nil {
		slog.Error("SaveDetailed: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("SaveDetailed: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(SaveDetailedResponse)

	if resp.Msg != "success" {
		slog.Error("SaveDetailed: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("SaveDetailed: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
