package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type ActivateListCalendarResponse struct {
	Today struct {
		Date  string `json:"date"`
		Week  string `json:"week"`
		Year  string `json:"year"`
		Month string `json:"month"`
		Day   string `json:"day"`
	} `json:"today"`
	Data []struct {
		Date   string `json:"date"`
		Day    string `json:"day"`
		Week   string `json:"week"`
		War    string `json:"war,omitempty"`
		Battle string `json:"battle,omitempty"`
		Orecar string `json:"orecar,omitempty"`
		School string `json:"school,omitempty"`
		Rescue string `json:"rescue,omitempty"`
	} `json:"data"`
}

func (c *Client) ActivateListCalendar(ctx context.Context, num int) (*ActivateListCalendarResponse, error) {
	params := &struct {
		Num int `json:"num"`
	}{
		Num: num,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("ActivateListCalendar: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/active/list/calendar", body)
	if err != nil {
		slog.Error("ActivateListCalendar: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("ActivateListCalendar: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(ActivateListCalendarResponse)

	if resp.Msg != "success" {
		slog.Error("ActivateListCalendar: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("ActivateListCalendar: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
