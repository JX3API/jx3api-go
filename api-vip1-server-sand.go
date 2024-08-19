package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type SandResponse struct {
	Zone   string `json:"zone"`
	Server string `json:"server"`
	Reset  int    `json:"reset"`
	Update int    `json:"update"`
	Data   []struct {
		TongId     int    `json:"tongId"`
		TongName   string `json:"tongName"`
		CastleId   int    `json:"castleId"`
		CastleName string `json:"castleName"`
		MasterId   int    `json:"masterId"`
		MasterName string `json:"masterName"`
		CampId     int    `json:"campId"`
		CampName   string `json:"campName"`
	} `json:"data"`
}

func (c *Client) ServerSand(ctx context.Context, server string) (*SandResponse, error) {
	params := &struct {
		Server string `json:"server"`
	}{
		Server: server,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("ServerSand: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/server/sand", body)
	if err != nil {
		slog.Error("ServerSand: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("ServerSand: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(SandResponse)

	if resp.Msg != "success" {
		slog.Error("ServerSand: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("ServerSand: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
