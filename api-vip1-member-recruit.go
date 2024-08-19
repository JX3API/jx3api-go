package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type MemberRecruitResponse struct {
	Zone   string `json:"zone"`
	Server string `json:"server"`
	Time   int64  `json:"time"`
	Data   []struct {
		CrossServer bool   `json:"crossServer"`
		ActivityId  int    `json:"activityId"`
		Activity    string `json:"activity"`
		Level       int    `json:"level"`
		Leader      string `json:"leader"`
		PushId      int    `json:"pushId"`
		RoomId      string `json:"roomID"`
		RoleId      int    `json:"roleId"`
		CreateTime  int64  `json:"createTime"`
		Number      int    `json:"number"`
		MaxNumber   int    `json:"maxNumber"`
		Label       []any  `json:"label"`
		Content     string `json:"content"`
	} `json:"data"`
}

func (c *Client) MemberRecruit(ctx context.Context, server string, keyword string, table int) (*MemberRecruitResponse, error) {
	params := &struct {
		Server  string `json:"server"`
		Keyword string `json:"keyword"`
		Table   int    `json:"table"`
	}{
		Server:  server,
		Keyword: keyword,
		Table:   table,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("MemberRecruit: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/member/recruit", body)
	if err != nil {
		slog.Error("MemberRecruit: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("MemberRecruit: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(MemberRecruitResponse)

	if resp.Msg != "success" {
		slog.Error("MemberRecruit: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("MemberRecruit: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
