package jx3api

import "context"

type RecruitSearchResponse struct {
	ActivityId         string `json:"activityId"`
	ActivityName       string `json:"activityName"`
	CampId             int    `json:"campId"`
	CampName           string `json:"campName"`
	CommentText        string `json:"commentText"`
	CreateTime         int64  `json:"createTime"`
	CurrentMemberCount int    `json:"currentMemberCount"`
	FlagBits           int    `json:"flagBits"`
	ForceId            int    `json:"forceId"`
	KungfuId           int    `json:"kungfuId"`
	LabelIdComment     string `json:"labelIdComment"`
	LabelIdTable       []int  `json:"labelIdTable"`
	LastModifyTime     int64  `json:"lastModifyTime"`
	MaxMemberCount     int    `json:"maxMemberCount"`
	MiniAvatarId       int    `json:"miniAvatarId"`
	RoleLevel          int    `json:"roleLevel"`
	RoleName           string `json:"roleName"`
	RoleType           int    `json:"roleType"`
	SubTypeName        string `json:"subTypeName"`
	TypeName           string `json:"typeName"`
	VariantInfo        struct {
		IsSwitchServer bool `json:"isSwitchServer"`
		PublisherId    int  `json:"publisherId"`
		PushId         int  `json:"pushId"`
		RoleId         int  `json:"roleId"`
	} `json:"variantInfo"`
}

// RecruitSearch 查询服务器的团队招募信息。
func (c *Client) RecruitSearch(ctx context.Context, server string, keyword string, typ int, limit int, token string) (*[]RecruitSearchResponse, error) {
	params := &struct {
		Server  string `json:"server"`
		Keyword string `json:"keyword,omitempty"`
		Label   int    `json:"label"`
		Limit   int    `json:"limit"`
		Token   string `json:"token,omitempty"`
	}{
		Server:  server,
		Keyword: keyword,
		Label:   typ,
		Limit:   limit,
		Token:   token,
	}
	return requestData[[]RecruitSearchResponse](ctx, c, "/recruit/search", params)
}
