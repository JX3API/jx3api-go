package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type SchoolToxicResponse struct {
	ID        int    `json:"id"`
	Class     string `json:"class"`
	Name      string `json:"name"`
	Toxic     string `json:"toxic"`
	Attribute string `json:"attribute"`
	Status    int    `json:"status"`
	Datetime  string `json:"datetime"`
}

func (c *Client) SchoolToxic(ctx context.Context, limit int) (*[]SchoolToxicResponse, error) {
	params := struct {
		Limit int `json:"limit"`
	}{
		Limit: limit,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("SchoolToxic: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/school/toxic", body)
	if err != nil {
		slog.Error("SchoolToxic: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("SchoolToxic: response body unmarshal error: ", err)
		return nil, err
	}

	data := new([]SchoolToxicResponse)

	if resp.Msg != "success" {
		slog.Error("SchoolToxic: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("SchoolToxic: data unmarshal error: ", err)
		return nil, err
	}

	return data, nil
}
