package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type SchoolMatrixResponse struct {
	Name      string `json:"name"`
	SkillName string `json:"skillName"`
	Descs     []struct {
		Desc  string `json:"desc"`
		Level int    `json:"level"`
		Name  string `json:"name"`
	} `json:"descs"`
}

func (c *Client) SchoolMatrix(ctx context.Context, name string) (*SchoolMatrixResponse, error) {
	params := &struct {
		Name string `json:"name"`
	}{
		Name: name,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("Matrix: request body marshal error: " + err.Error())
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
		slog.Info("Matrix: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(SchoolMatrixResponse)

	if resp.Msg != "success" {
		slog.Error("Matrix: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("Matrix: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
