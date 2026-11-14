package jx3api

import "context"

type VoiceChannelsResponse struct {
	Data []struct {
		Asid     int    `json:"asid"`
		CampName string `json:"campName"`
		Esid     string `json:"esid"`
		Limit    int    `json:"limit"`
		Logo     int    `json:"logo"`
		LogoUrl  string `json:"logoUrl"`
		Sid      int    `json:"sid"`
		Snick    string `json:"snick"`
		Users    int    `json:"users"`
	} `json:"data"`
	Server string `json:"server"`
}

type DuowanStatisticsServer = VoiceChannelsResponse

// VoiceChannels 查询统战 YY 频道的在线主播。
func (c *Client) VoiceChannels(ctx context.Context, server string) (*[]VoiceChannelsResponse, error) {
	params := &struct {
		Server string `json:"server,omitempty"`
	}{
		Server: server,
	}
	return requestData[[]VoiceChannelsResponse](ctx, c, "/voice/channels", params)
}

// DuowanStatistics is an alias for VoiceChannels.
func (c *Client) DuowanStatistics(ctx context.Context, server string, token string) (*[]VoiceChannelsResponse, error) {
	return c.VoiceChannels(ctx, server)
}
