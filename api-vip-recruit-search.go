package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type RecruitData struct {
	Activity    string   `json:"activity"`
	Leader      string   `json:"leader"`
	Content     string   `json:"content"`
	CrossServer bool     `json:"crossServer"`
	ActivityId  int      `json:"activityId"`
	Level       int      `json:"level"`
	PushId      int      `json:"pushId"`
	RoomID      string   `json:"roomID"`
	RoleId      int      `json:"roleId"`
	CreateTime  int64    `json:"createTime"`
	Number      int      `json:"number"`
	MaxNumber   int      `json:"maxNumber"`
	Label       []string `json:"label"`
}

type RecruitSearchResponse struct {
	Zone   string        `json:"zone"`
	Server string        `json:"server"`
	Type   int           `json:"type"`
	Data   []RecruitData `json:"data"`
	Time   int64         `json:"time"`
}

func (c *Client) RecruitSearch(ctx context.Context, server string, keyword string, typ int, token string) (*RecruitSearchResponse, error) {
	params := &struct {
		Server  string `json:"server"`
		Keyword string `json:"keyword"`
		Type    int    `json:"type"`
		Token   string `json:"token"`
	}{
		Server:  server,
		Keyword: keyword,
		Type:    typ,
		Token:   token,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("RecruitSearch: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/recruit/search", body)
	if err != nil {
		slog.Error("RecruitSearch: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("RecruitSearch: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(RecruitSearchResponse)

	if resp.Msg != "success" {
		slog.Error("RecruitSearch: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("RecruitSearch: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
