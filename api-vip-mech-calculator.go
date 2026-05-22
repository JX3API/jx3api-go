package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type MechCalculatorResponse struct {
	NowTime     string `json:"now_time"`
	NowNode     string `json:"now_node"`
	NowResult   string `json:"now_result"`
	NextNode    string `json:"next_node"`
	NextResult  string `json:"next_result"`
	IntervalTime string `json:"interval_time"`
}

func (c *Client) MechCalculator(ctx context.Context) (*MechCalculatorResponse, error) {
	params := &struct{}{}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("MechCalculator: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/mech/calculator", body)
	if err != nil {
		slog.Error("MechCalculator: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("MechCalculator: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(MechCalculatorResponse)

	if resp.Msg != "success" {
		slog.Error("MechCalculator: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("MechCalculator: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
