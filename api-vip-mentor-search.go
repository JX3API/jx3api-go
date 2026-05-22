package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type MentorData struct {
	RoleId         int64  `json:"roleId"`
	RoleName       string `json:"roleName"`
	RoleLevel      int    `json:"roleLevel"`
	CampName       string `json:"campName"`
	TongName       string `json:"tongName"`
	TongMasterName string `json:"tongMasterName"`
	BodyId         int    `json:"bodyId"`
	BodyName       string `json:"bodyName"`
	ForceId        int    `json:"forceId"`
	ForceName      string `json:"forceName"`
	Comment        string `json:"comment"`
}

type MentorSearchResponse struct {
	Zone   string       `json:"zone"`
	Server string       `json:"server"`
	Type   int          `json:"type"`
	Data   []MentorData `json:"data"`
	Time   int64        `json:"time"`
}

func (c *Client) MentorSearch(ctx context.Context, typ int, server string, keyword string, token string) (*MentorSearchResponse, error) {
	params := &struct {
		Type    int    `json:"type"`
		Server  string `json:"server"`
		Keyword string `json:"keyword"`
		Token   string `json:"token"`
	}{
		Type:    typ,
		Server:  server,
		Keyword: keyword,
		Token:   token,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("MentorSearch: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/mentor/search", body)
	if err != nil {
		slog.Error("MentorSearch: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("MentorSearch: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(MentorSearchResponse)

	if resp.Msg != "success" {
		slog.Error("MentorSearch: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("MentorSearch: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
