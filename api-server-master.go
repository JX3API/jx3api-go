package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type ServerMasterResponse struct {
	ID     int    `json:"id"`
	Zone   string `json:"zone"`
	Server string `json:"server"`
	Status int    `json:"status"`
	Time   int    `json:"time"`
}

func (c *Client) ServerMaster(ctx context.Context, name string) (*ServerMasterResponse, error) {
	params := struct {
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
		slog.Info("ServerMaster: response body unmarshal error: ", err)
		return nil, err
	}

	data := new(ServerMasterResponse)

	if resp.Msg != "success" {
		slog.Error("ServerMaster: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("ServerMaster: data unmarshal error: ", err)
		return nil, err
	}

	return data, nil
}
