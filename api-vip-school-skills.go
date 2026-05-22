package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type SkillItem struct {
	Name        string `json:"name"`
	SimpleDesc  string `json:"simpleDesc"`
	Desc        string `json:"desc"`
	SpecialDesc string `json:"specialDesc"`
	Interval    string `json:"interval"`
	Consumption string `json:"consumption"`
	Distance    string `json:"distance"`
	Icon        string `json:"icon"`
	Kind        string `json:"kind"`
	SubKind     string `json:"subKind"`
	ReleaseType string `json:"releaseType"`
	Weapon      string `json:"weapon"`
}

type SkillGroup struct {
	Class string      `json:"class"`
	Data  []SkillItem `json:"data"`
}

type SchoolSkillsResponse struct {
	Data []SkillGroup `json:"data"`
}

func (c *Client) SchoolSkills(ctx context.Context, name string, ticket string, token string) (*SchoolSkillsResponse, error) {
	params := &struct {
		Name   string `json:"name"`
		Ticket string `json:"ticket"`
		Token  string `json:"token"`
	}{
		Name:   name,
		Ticket: ticket,
		Token:  token,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("SchoolSkills: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/school/skills", body)
	if err != nil {
		slog.Error("Skills: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("SchoolSkills: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(SchoolSkillsResponse)

	if resp.Msg != "success" {
		slog.Error("SchoolSkills: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("SchoolSkills: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
