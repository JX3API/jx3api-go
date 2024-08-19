package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type ServerMasterResponse struct {
	ID           string   `json:"id"`
	Zone         string   `json:"zone"`
	Name         string   `json:"name"`
	Column       string   `json:"column"`
	Duowan       Duowan   `json:"duowan"`
	Abbreviation []string `json:"abbreviation"`
	Subordinate  []string `json:"subordinate"`
}

type Duowan struct {
	HaoqiMeng []int `json:"浩气盟"`
	ErenGu    []int `json:"恶人谷"`
}

func (c *Client) ServerMaster(ctx context.Context, name string) (*ServerMasterResponse, error) {
	params := &struct {
		Name string `json:"name"`
	}{
		Name: name,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("ServerMaster: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/server/master", body)
	if err != nil {
		slog.Error("ServerMaster: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("ServerMaster: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new(ServerMasterResponse)

	if resp.Msg != "success" {
		slog.Error("ServerMaster: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("ServerMaster: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
