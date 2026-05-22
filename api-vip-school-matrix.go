package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type SchoolMatrixData struct {
	Desc  string `json:"desc"`
	Level int    `json:"level"`
	Name  string `json:"name"`
}

type SchoolMatrixResponse struct {
	Name      string            `json:"name"`
	SkillName string            `json:"skillName"`
	Data      []SchoolMatrixData `json:"data"`
}

func (c *Client) SchoolMatrix(ctx context.Context, name string, ticket string, token string) (*SchoolMatrixResponse, error) {
	params := &struct {
		Name   string `json:"name"`
		Ticket string `json:"ticket"`
		Token  string `json:"token"`
	}{
		Name:   name,
		Ticket: ticket,
		Token:  token,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("SchoolMatrix: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/school/matrix", body)
	if err != nil {
		slog.Error("Matrix: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("SchoolMatrix: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(SchoolMatrixResponse)

	if resp.Msg != "success" {
		slog.Error("SchoolMatrix: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("SchoolMatrix: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
