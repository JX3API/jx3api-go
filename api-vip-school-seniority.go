package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type SeniorityResponse struct {
	ZoneName   string `json:"zoneName"`
	ServerName string `json:"serverName"`
	RoleName   string `json:"roleName"`
	RoleId     int    `json:"roleId"`
	ForceName  string `json:"forceName"`
	ForceId    int    `json:"forceId"`
	ForceIcon  string `json:"forceIcon"`
	AvatarUrl  string `json:"avatarUrl"`
	Seniority  int    `json:"seniority"`
}

func (c *Client) SchoolSeniority(ctx context.Context, server string, school string, ticket string, token string) (*[]SeniorityResponse, error) {
	params := &struct {
		Server string `json:"server"`
		School string `json:"school"`
		Ticket string `json:"ticket"`
		Token  string `json:"token"`
	}{
		Server: server,
		School: school,
		Ticket: ticket,
		Token:  token,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("SchoolSeniority: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/school/seniority", body)
	if err != nil {
		slog.Error("SchoolSeniority: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("SchoolSeniority: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]SeniorityResponse)

	if resp.Msg != "success" {
		slog.Error("SchoolSeniority: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("SchoolSeniority: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
