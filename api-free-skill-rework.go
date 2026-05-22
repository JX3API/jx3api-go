package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type SkillReworkResponse struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
	Time  string `json:"time"`
}

func (c *Client) SkillRework(ctx context.Context) (*[]SkillReworkResponse, error) {
	params := &struct{}{}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("SkillRework: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/skill/rework", body)
	if err != nil {
		slog.Error("SkillRework: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("SkillRework: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]SkillReworkResponse)

	if resp.Msg != "success" {
		slog.Error("SkillRework: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("SkillRework: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
