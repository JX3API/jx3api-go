package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type ArenaSchoolsResponse struct {
	Name string `json:"name"`
	This int    `json:"this"`
	Last int    `json:"last"`
}

func (c *Client) ArenaSchools(ctx context.Context, mode int, ticket string, token string) (*[]ArenaSchoolsResponse, error) {
	params := &struct {
		Mode   int    `json:"mode"`
		Ticket string `json:"ticket"`
		Token  string `json:"token"`
	}{
		Mode:   mode,
		Ticket: ticket,
		Token:  token,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("ArenaSchools: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/arena/schools", body)
	if err != nil {
		slog.Error("ArenaSchools: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("ArenaSchools: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]ArenaSchoolsResponse)

	if resp.Msg != "success" {
		slog.Error("ArenaSchools: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("ArenaSchools: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
