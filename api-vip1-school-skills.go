package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type SkillResponse struct {
	Class string `json:"class"`
	Data  []struct {
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
	} `json:"data"`
	Time int `json:"time"`
}

func (c *Client) SchoolSkills(ctx context.Context, name string) (*[]SkillResponse, error) {
	params := &struct {
		Name string `json:"name"`
	}{
		Name: name,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("Skills: request body marshal error: " + err.Error())
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
		slog.Info("Skills: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]SkillResponse)

	if resp.Msg != "success" {
		slog.Error("Skills: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("Skills: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
