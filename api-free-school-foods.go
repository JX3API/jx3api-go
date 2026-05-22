package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type SchoolFoodsResponse struct {
	ID     int    `json:"id"`
	School string `json:"school"`
	Kungfu string `json:"kungfu"`
	Color  string `json:"color"`
	Class  string `json:"class"`
	Name   string `json:"name"`
	Boost  string `json:"boost"`
}

func (c *Client) SchoolFoods(ctx context.Context) (*[]SchoolFoodsResponse, error) {
	params := &struct{}{}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("SchoolFoods: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/school/foods", body)
	if err != nil {
		slog.Error("SchoolFoods: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("SchoolFoods: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]SchoolFoodsResponse)

	if resp.Msg != "success" {
		slog.Error("SchoolFoods: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("SchoolFoods: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
