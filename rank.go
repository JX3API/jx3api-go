package jx3api

import "encoding/json"

type RankServer struct {
	Zone   string                       `json:"zone"`
	Server string                       `json:"server"`
	Label  string                       `json:"label"`
	Data   []map[string]json.RawMessage `json:"data"`
	Time   int64                        `json:"time"`
}

// RankResponse contains one server when server is specified, otherwise all servers.
type RankResponse []RankServer

func (r *RankResponse) UnmarshalJSON(raw []byte) error {
	return unmarshalRecords(raw, (*[]RankServer)(r))
}
