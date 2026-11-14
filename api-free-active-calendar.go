package jx3api

import "context"

type ActiveCalendarResponse struct {
	Battle string   `json:"battle"`
	Card   []string `json:"card"`
	Date   string   `json:"date"`
	Draw   []string `json:"draw"`
	Leader []string `json:"leader,omitempty"`
	Lucky  []string `json:"lucky"`
	Orecar string   `json:"orecar"`
	Rescue string   `json:"rescue"`
	School string   `json:"school"`
	War    string   `json:"war"`
	Week   string   `json:"week"`
	Weekly struct {
		Conn []string `json:"conn"`
		Raid []string `json:"raid"`
	} `json:"weekly"`
}

type ActivateCalendarResponse = ActiveCalendarResponse

// ActiveCalendar 查询每日活动安排及本周秘境、副本信息。
func (c *Client) ActiveCalendar(ctx context.Context, num int) (*ActiveCalendarResponse, error) {
	params := &struct {
		Num  int    `json:"num"`
		Mode string `json:"mode"`
	}{
		Num:  num,
		Mode: "day",
	}
	return requestData[ActiveCalendarResponse](ctx, c, "/active/calendar", params)
}

// ActivateCalendar is an alias for ActiveCalendar.
func (c *Client) ActivateCalendar(ctx context.Context, server string, num int) (*ActiveCalendarResponse, error) {
	return c.ActiveCalendar(ctx, num)
}
