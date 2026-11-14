package jx3api

import "context"

type SandRecordsResponse struct {
	Data []struct {
		CampId      int    `json:"campId"`
		CampName    string `json:"campName"`
		CastleId    int    `json:"castleId"`
		CastleName  string `json:"castleName"`
		DefendCount int    `json:"defendCount"`
		MasterId    int    `json:"masterId"`
		MasterName  string `json:"masterName"`
		RidePiece   int    `json:"ridePiece"`
		TongId      int    `json:"tongId"`
		TongName    string `json:"tongName"`
	} `json:"data"`
	Reset  int64  `json:"reset"`
	Server string `json:"server"`
	Update int64  `json:"update"`
	Zone   string `json:"zone"`
}

// SandRecords 查询据点归属及本周祭祀次数。
func (c *Client) SandRecords(ctx context.Context, server string, token string) (*SandRecordsResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Token:  token,
	}
	return requestData[SandRecordsResponse](ctx, c, "/sand/records", params)
}
