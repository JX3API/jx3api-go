package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type AchievementResponse struct {
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
		ID           int      `json:"id"`
		Icon         string   `json:"icon"`
		Likes        int      `json:"likes"`
		Name         string   `json:"name"`
		Class        string   `json:"class"`
		SubClass     string   `json:"subClass"`
		Desc         string   `json:"desc"`
		Detail       string   `json:"detail"`
		Maps         []int    `json:"maps"`
		IsFinished   bool     `json:"isFinished"`
		IsFav        bool     `json:"isFav"`
		Type         string   `json:"type"`
		CurrentValue int      `json:"currentValue"`
		TriggerValue int      `json:"triggerValue"`
		Subset       []int    `json:"subset"`
		RewardItem   struct{} `json:"rewardItem"`
		RewardPoint  int      `json:"rewardPoint"`
		RewardPrefix string   `json:"rewardPrefix"`
		RewardSuffix string   `json:"rewardSuffix"`
	} `json:"data"`
}

func (c *Client) RoleAchievement(ctx context.Context, server string, role string, name string) (*AchievementResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Role   string `json:"role"`
		Name   string `json:"name"`
		Ticket string `json:"ticket"`
	}{
		Server: server,
		Role:   role,
		Name:   name,
		Ticket: c.Opts.Ticket,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("RoleAchievement: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/role/achievement", body)
	if err != nil {
		slog.Error("RoleAchievement: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("RoleAchievement: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(AchievementResponse)

	if resp.Msg != "success" {
		slog.Error("RoleAchievement: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("RoleAchievement: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
