package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type TiebaRandomResponse struct {
	ID     int    `json:"id"`
	Class  string `json:"class"`
	Zone   string `json:"zone"`
	Server string `json:"server"`
	Name   string `json:"name"`
	Title  string `json:"title"`
	URL    int64  `json:"url"`
	Date   string `json:"date"`
}

func (c *Client) TiebaRandom(ctx context.Context, class string, server string, limit int) (*[]TiebaRandomResponse, error) {
	params := &struct {
		Class  string `json:"class"`
		Server string `json:"server"`
		Limit  int    `json:"limit"`
	}{
		Class:  class,
		Server: server,
		Limit:  limit,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("TiebaRandom: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/tieba/random", body)
	if err != nil {
		slog.Error("TiebaRandom: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("TiebaRandom: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]TiebaRandomResponse)

	if resp.Msg != "success" {
		slog.Error("TiebaRandom: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("TiebaRandom: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
