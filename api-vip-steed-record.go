package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type SteedRecordResponse struct {
	ID              int    `json:"id"`
	Zone            string `json:"zone"`
	Server          string `json:"server"`
	MapName         string `json:"map_name"`
	RefreshTime     int64  `json:"refresh_time"`
	CaptureRoleName string `json:"capture_role_name"`
	CaptureCampName string `json:"capture_camp_name"`
	CaptureTime     int64  `json:"capture_time"`
	AuctionRoleName string `json:"auction_role_name"`
	AuctionCampName string `json:"auction_camp_name"`
	AuctionTime     int64  `json:"auction_time"`
	AuctionAmount   string `json:"auction_amount"`
	StartTime       int64  `json:"start_time"`
	EndTime         int64  `json:"end_time"`
}

func (c *Client) SteedRecord(ctx context.Context, server string, token string) (*[]SteedRecordResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Token  string `json:"token"`
	}{
		Server: server,
		Token:  token,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("SteedRecord: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/steed/records", body)
	if err != nil {
		slog.Error("SteedRecord: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("SteedRecord: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]SteedRecordResponse)

	if resp.Msg != "success" {
		slog.Error("SteedRecord: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("SteedRecord: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
