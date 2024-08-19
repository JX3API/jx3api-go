package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type HorseRecordResponse struct {
	ID              int    `json:"id"`
	Zone            string `json:"zone"`
	Server          string `json:"server"`
	Name            string `json:"name"`
	Level           int    `json:"level"`
	MapName         string `json:"map_name"`
	RefreshTime     int    `json:"refresh_time"`
	CaptureRoleName string `json:"capture_role_name"`
	CaptureCampName string `json:"capture_camp_name"`
	CaptureTime     int    `json:"capture_time"`
	AuctionRoleName string `json:"auction_role_name"`
	AuctionCampName string `json:"auction_camp_name"`
	AuctionTime     int    `json:"auction_time"`
	AuctionAmount   string `json:"auction_amount"`
	StartTime       int    `json:"start_time"`
	EndTime         int    `json:"end_time"`
}

func (c *Client) HorseRecord(ctx context.Context, server string) (*[]HorseRecordResponse, error) {
	params := &struct {
		Server string `json:"server"`
	}{
		Server: server,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("HorseRecord: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/horse/record", body)
	if err != nil {
		slog.Error("HorseRecord: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("HorseRecord: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]HorseRecordResponse)

	if resp.Msg != "success" {
		slog.Error("HorseRecord: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("HorseRecord: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
