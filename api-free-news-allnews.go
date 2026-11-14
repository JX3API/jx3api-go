package jx3api

import (
	"context"
	"encoding/json"
)

type NewsRecordsResponse struct {
	Catid string `json:"catid"`
	Date  string `json:"date"`
	Desc  struct {
		AppAllowComment  *string         `json:"app_allow_comment,omitempty"`
		AppIsRecommended *string         `json:"app_is_recommended,omitempty"`
		Asktime          *string         `json:"asktime,omitempty"`
		AuthorLive       *string         `json:"author_live,omitempty"`
		AuthorLiveIcon   *string         `json:"author_live_icon,omitempty"`
		Catid            *string         `json:"catid,omitempty"`
		CheckIsLive      *string         `json:"check_is_live,omitempty"`
		Content          string          `json:"content"`
		Description      *string         `json:"description,omitempty"`
		ExtData          json.RawMessage `json:"ext_data,omitempty"`
		Id               *string         `json:"id,omitempty"`
		Inputtime        json.RawMessage `json:"inputtime"`
		Islink           *string         `json:"islink,omitempty"`
		Keywords         *string         `json:"keywords,omitempty"`
		Listorder        *string         `json:"listorder,omitempty"`
		LiveAuthor       *string         `json:"live_author,omitempty"`
		LiveAuthorIcon   *string         `json:"live_author_icon,omitempty"`
		LiveStatus       *string         `json:"live_status,omitempty"`
		Oid              *string         `json:"oid,omitempty"`
		Style            *string         `json:"style,omitempty"`
		Thumb            *string         `json:"thumb,omitempty"`
		Title            string          `json:"title"`
		Top              *string         `json:"top,omitempty"`
		Updatetime       *string         `json:"updatetime,omitempty"`
		Url              *string         `json:"url,omitempty"`
	} `json:"desc"`
	Title string `json:"title"`
	Type  string `json:"type"`
	Url   string `json:"url"`
}

type NewsAllnewsResponse = NewsRecordsResponse

// NewsRecords 查询最新游戏新闻资讯。
func (c *Client) NewsRecords(ctx context.Context, limit int) (*[]NewsRecordsResponse, error) {
	params := &struct {
		Limit int `json:"limit"`
	}{
		Limit: limit,
	}
	return requestData[[]NewsRecordsResponse](ctx, c, "/news/records", params)
}

// NewsAllnews is an alias for NewsRecords.
func (c *Client) NewsAllnews(ctx context.Context, limit int) (*[]NewsRecordsResponse, error) {
	return c.NewsRecords(ctx, limit)
}
