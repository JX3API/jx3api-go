package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type ServerEventResponse struct {
	ID                int    `json:"id"`
	CampName          string `json:"camp_name"`
	FenxianZoneName   string `json:"fenxian_zone_name"`
	FenxianServerName string `json:"fenxian_server_name"`
	FriendZoneName    string `json:"friend_zone_name"`
	FriendServerName  string `json:"friend_server_name"`
	RoleName          string `json:"role_name"`
	AddTime           int64  `json:"add_time"`
}

func (c *Client) ServerEvent(ctx context.Context, name string, limit int) (*[]ServerEventResponse, error) {
	params := &struct {
		Name  string `json:"name"`
		Limit int    `json:"limit"`
	}{
		Name:  name,
		Limit: limit,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("ServerEvent: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/server/event", body)
	if err != nil {
		slog.Error("ServerEvent: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("ServerEvent: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]ServerEventResponse)

	if resp.Msg != "success" {
		slog.Error("ServerEvent: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("ServerEvent: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
