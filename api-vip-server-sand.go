package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type SandRecord struct {
	TongId     int    `json:"tongId"`
	TongName   string `json:"tongName"`
	CastleId   int    `json:"castleId"`
	CastleName string `json:"castleName"`
	MasterId   int    `json:"masterId"`
	MasterName string `json:"masterName"`
	CampId     int    `json:"campId"`
	CampName   string `json:"campName"`
}

type SandRecordsResponse struct {
	Zone   string       `json:"zone"`
	Server string       `json:"server"`
	Reset  int          `json:"reset"`
	Update int          `json:"update"`
	Data   []SandRecord `json:"data"`
}

func (c *Client) SandRecords(ctx context.Context, server string) (*SandRecordsResponse, error) {
	params := &struct {
		Server string `json:"server"`
	}{
		Server: server,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("SandRecords: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/sand/records", body)
	if err != nil {
		slog.Error("ServerSand: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("SandRecords: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(SandRecordsResponse)

	if resp.Msg != "success" {
		slog.Error("SandRecords: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("SandRecords: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
