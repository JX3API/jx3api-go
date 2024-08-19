package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type SchoolForceResponse struct {
	Level int `json:"level"`
	Data  []struct {
		Name    string `json:"name"`
		Class   int    `json:"class"`
		Desc    string `json:"desc"`
		Icon    string `json:"icon"`
		Kind    string `json:"kind"`
		SubKind string `json:"subKind"`
	} `json:"data"`
}

func (c *Client) SchoolForce(ctx context.Context, name string) (*[]SchoolForceResponse, error) {
	params := &struct {
		Name string `json:"name"`
	}{
		Name: name,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("Force: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/school/force", body)
	if err != nil {
		slog.Error("Force: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("Force: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]SchoolForceResponse)

	if resp.Msg != "success" {
		slog.Error("Force: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("Force: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
