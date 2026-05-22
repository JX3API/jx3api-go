package jx3api

import "encoding/json"

type Response struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg"`
	Data *json.RawMessage `json:"data"`
	Time int              `json:"time"`
}
