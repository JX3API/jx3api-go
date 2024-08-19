package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type StudentListResponse struct {
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

func (c *Client) MemberStudent(ctx context.Context, server string, keyword *string) (*StudentListResponse, error) {
	params := &struct {
		Server  string  `json:"server"`
		Keyword *string `json:"keyword,omitempty"`
	}{
		Server:  server,
		Keyword: keyword,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("MemberStudent: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/member/student", body)
	if err != nil {
		slog.Error("MemberStudent: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("MemberStudent: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(StudentListResponse)

	if resp.Msg != "success" {
		slog.Error("MemberStudent: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("MemberStudent: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
