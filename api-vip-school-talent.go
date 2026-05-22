package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type TalentData struct {
	Name          string `json:"name"`
	Class         int    `json:"class"`
	SimpleDesc    string `json:"simpleDesc"`
	Desc          string `json:"desc"`
	SpecialDesc   string `json:"specialDesc"`
	Interval      string `json:"interval"`
	Consumption   string `json:"consumption"`
	Distance      string `json:"distance"`
	Icon          string `json:"icon"`
	Kind          string `json:"kind"`
	SubKind       string `json:"subKind"`
	ReleaseType   string `json:"releaseType"`
	Weapon        string `json:"weapon"`
}

type TalentLevel struct {
	Level int            `json:"level"`
	Data  []TalentData   `json:"data"`
}

type SchoolTalentResponse struct {
	Data []TalentLevel `json:"data"`
}

func (c *Client) SchoolTalent(ctx context.Context, name string, ticket string, token string) (*[]TalentLevel, error) {
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
		slog.Error("SchoolTalent: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/school/talent", body)
	if err != nil {
		slog.Error("SchoolTalent: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("SchoolTalent: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]TalentLevel)

	if resp.Msg != "success" {
		slog.Error("SchoolTalent: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("SchoolTalent: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
