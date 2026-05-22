package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type RoleMonsterSkill struct {
	SkillCost    int    `json:"skill_cost"`
	SkillName    string `json:"skill_name"`
	LeaderName   string `json:"leader_name"`
	SkillColor   int    `json:"skill_color"`
	SkillLevel   int    `json:"skill_level"`
	IsDeprecated bool   `json:"is_deprecated"`
}

type RoleMonsterResponse struct {
	Zone        string             `json:"zone"`
	Server      string             `json:"server"`
	RoleName    string             `json:"role_name"`
	RoleId      string             `json:"role_id"`
	GlobalId    string             `json:"global_id"`
	SkillStamina int                `json:"skill_stamina"`
	SkillEnergy int                 `json:"skill_energy"`
	SkillCount  int                `json:"skill_count"`
	SkillList   []RoleMonsterSkill `json:"skill_list"`
	UpdateTime  int64              `json:"update_time"`
}

func (c *Client) RoleMonster(ctx context.Context, server string, name string, token string) (*RoleMonsterResponse, error) {
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
		slog.Error("RoleMonster: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/role/monster", body)
	if err != nil {
		slog.Error("RoleMonster: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("RoleMonster: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(RoleMonsterResponse)

	if resp.Msg != "success" {
		slog.Error("RoleMonster: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("RoleMonster: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
