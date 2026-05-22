package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type FenxianRecord struct {
	ID          int    `json:"id"`
	CampName    string `json:"camp_name"`
	FenxianName string `json:"fenxian_name"`
	FriendName  string `json:"friend_name"`
	RoleName    string `json:"role_name"`
	SeizeTime   int64  `json:"seize_time"`
}

func (c *Client) FenxianRecords(ctx context.Context) (*[]FenxianRecord, error) {
	raw, err := c.request(ctx, "/data/fenxian/records", nil)
	if err != nil {
		slog.Error("FenxianRecords: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("FenxianRecords: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]FenxianRecord)

	if resp.Msg != "success" {
		slog.Error("FenxianRecords: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("FenxianRecords: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
