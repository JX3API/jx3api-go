package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type TeacherListResponse struct {
	Zone   string `json:"zone"`
	Server string `json:"server"`
	Data   []struct {
		RoleId         int    `json:"roleId"`
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
		Time           int64  `json:"time"`
	} `json:"data"`
}

func (c *Client) MemberTeacher(ctx context.Context, server string, keyword *string) (*TeacherListResponse, error) {
	params := &struct {
		Server  string  `json:"server"`
		Keyword *string `json:"keyword,omitempty"`
	}{
		Server:  server,
		Keyword: keyword,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("MemberTeacher: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/member/teacher", body)
	if err != nil {
		slog.Error("MemberTeacher: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("MemberTeacher: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(TeacherListResponse)

	if resp.Msg != "success" {
		slog.Error("MemberTeacher: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("MemberTeacher: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
