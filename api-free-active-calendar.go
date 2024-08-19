package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type ActivateCalendarResponse struct {
	Date   string   `json:"date"`
	Week   string   `json:"week"`
	War    string   `json:"war"`
	Battle string   `json:"battle"`
	Orecar string   `json:"orecar"`
	School string   `json:"school"`
	Rescue string   `json:"rescue"`
	Draw   string   `json:"draw,omitempty"`
	Leader []string `json:"leader,omitempty"`
	Team   []string `json:"team"`
}

func (c *Client) ActivateCalendar(ctx context.Context, server string, num int) (*ActivateCalendarResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Num    int    `json:"num"`
	}{
		Server: server,
		Num:    num,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("ActivateCalendar: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/active/calendar", body)
	if err != nil {
		slog.Error("ActivateCalendar: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("ActivateCalendar: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(ActivateCalendarResponse)

	if resp.Msg != "success" {
		slog.Error("ActivateCalendar: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("ActivateCalendar: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
