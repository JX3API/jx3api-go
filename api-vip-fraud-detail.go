package jx3api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
)

type FraudData struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	Tid   int64  `json:"tid"`
	Text  string `json:"text"`
	Time  int64  `json:"time"`
}

type FraudDetailResponse struct {
	Server string       `json:"server"`
	Tieba  string       `json:"tieba"`
	Data   []FraudData  `json:"data"`
}

func (c *Client) FraudDetail(ctx context.Context, uid int, token string) (*[]FraudDetailResponse, error) {
	params := &struct {
		Uid   int    `json:"uid"`
		Token string `json:"token"`
	}{
		Uid:   uid,
		Token: token,
	}

	buf, err := json.Marshal(params)
	if err != nil {
		slog.Error("FraudDetail: request body marshal error: " + err.Error())
		return nil, err
	}

	body := bytes.NewReader(buf)

	raw, err := c.request(ctx, "/data/fraud/detail", body)
	if err != nil {
		slog.Error("FraudDetail: request error: " + err.Error())
		return nil, err
	}

	resp := new(Response)

	if err := json.Unmarshal(raw, &resp); err != nil {
		slog.Info("FraudDetail: response body unmarshal error: " + err.Error())
		return nil, err
	}

	data := new([]FraudDetailResponse)

	if resp.Msg != "success" {
		slog.Error("FraudDetail: API error: " + resp.Msg)
		return nil, errors.New(resp.Msg)
	}

	if err := json.Unmarshal(*resp.Data, &data); err != nil {
		slog.Info("FraudDetail: data unmarshal error: " + err.Error())
		return nil, err
	}

	return data, nil
}
