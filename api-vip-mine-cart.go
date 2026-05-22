package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type MineCartData struct {
	ID        int    `json:"id"`
	Zone      string `json:"zone"`
	Server    string `json:"server"`
	Leader    string `json:"leader"`
	CampName  string `json:"camp_name"`
	Castle    string `json:"castle"`
	Status    int    `json:"status"`
	StrStatus string `json:"str_status"`
}

type MineCartServer struct {
	Server string          `json:"server"`
	Data   []MineCartData  `json:"data"`
}

func (c *Client) MineCart(ctx context.Context, token string) (*[]MineCartServer, error) {
	params := &struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("MineCart: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/mine/cart", body)
	if err != nil {
		slog.Error("MineCart: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("MineCart: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]MineCartServer)

	if resp.Msg != "success" {
		slog.Error("MineCart: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("MineCart: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
