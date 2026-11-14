package jx3api

import (
	"context"
)

type ActiveListCalendarResponse struct {
	Today struct {
		Date  string `json:"date"`
		Day   string `json:"day"`
		Month string `json:"month"`
		Week  string `json:"week"`
		Year  string `json:"year"`
	} `json:"today"`
	Total []struct {
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
	} `json:"total"`
}

// ActiveListCalendar 查询每日活动安排及本周秘境、副本信息。
func (c *Client) ActiveListCalendar(ctx context.Context, num int) (*ActiveListCalendarResponse, error) {
	params := &struct {
		Num  int    `json:"num"`
		Mode string `json:"mode"`
	}{
		Num:  num,
		Mode: "list",
	}
	return requestData[ActiveListCalendarResponse](ctx, c, "/active/calendar", params)
}

type ActivateListCalendarResponse = ActiveListCalendarResponse

// ActivateListCalendar is an alias for ActiveListCalendar.
func (c *Client) ActivateListCalendar(ctx context.Context, num int) (*ActiveListCalendarResponse, error) {
	return c.ActiveListCalendar(ctx, num)
}
