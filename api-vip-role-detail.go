package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type RoleDetailResponse struct {
	ZoneName   string `json:"zoneName"`
	ServerName string `json:"serverName"`
	RoleName   string `json:"roleName"`
	RoleId     string `json:"roleId"`
	GlobalId   string `json:"globalId"`
	ForceName  string `json:"forceName"`
	ForceId    int    `json:"forceId"`
	BodyName   string `json:"bodyName"`
	BodyId     int    `json:"bodyId"`
	TongName   string `json:"tongName"`
	TongId     int    `json:"tongId"`
	CampName   string `json:"campName"`
	CampId     int    `json:"campId"`
}

func (c *Client) RoleDetail(ctx context.Context, server string, name string, token string) (*RoleDetailResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Name   string `json:"name"`
		Token  string `json:"token"`
	}{
		Server: server,
		Name:   name,
		Token:  token,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("RoleDetail: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/role/detail", body)
	if err != nil {
		slog.Error("RoleDetail: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("RoleDetail: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(RoleDetailResponse)

	if resp.Msg != "success" {
		slog.Error("RoleDetail: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("RoleDetail: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
