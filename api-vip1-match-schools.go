package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type MatchSchoolsResponse struct {
	Name string `json:"name"`
	This int    `json:"this"`
	Last int    `json:"last"`
}

func (c *Client) MatchSchools(ctx context.Context, mode int) (*[]MatchSchoolsResponse, error) {
	params := &struct {
		Mode   int    `json:"mode"`
		Ticket string `json:"ticket"`
	}{
		Mode:   mode,
		Ticket: c.Opts.Ticket,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("MatchSchools: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/match/schools", body)
	if err != nil {
		slog.Error("MatchSchools: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("MatchSchools: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]MatchSchoolsResponse)

	if resp.Msg != "success" {
		slog.Error("MatchSchools: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("MatchSchools: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
