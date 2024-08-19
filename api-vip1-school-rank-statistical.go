package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type SchoolRankStatisticalResponse struct {
	Name   string `json:"name"`
	Role   string `json:"role"`
	School string `json:"school"`
	Server string `json:"server"`
	Zone   string `json:"zone"`
	Value  int    `json:"value"`
	Avatar string `json:"avatar"`
}

func (c *Client) SchoolRankStatistical(ctx context.Context, school string, server string) (*[]SchoolRankStatisticalResponse, error) {
	params := &struct {
		School string `json:"school"`
		Server string `json:"server"`
		Ticket string `json:"ticket"`
	}{
		School: school,
		Server: server,
		Ticket: c.Opts.Ticket,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("SchoolRankStatistical: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/school/rank/statistical", body)
	if err != nil {
		slog.Error("SchoolRankStatistical: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("SchoolRankStatistical: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]SchoolRankStatisticalResponse)

	if resp.Msg != "success" {
		slog.Error("SchoolRankStatistical: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("SchoolRankStatistical: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
