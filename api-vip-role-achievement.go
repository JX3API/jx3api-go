package jx3api

import "context"

type RoleAchievementResponse struct {
	BodyId   int    `json:"bodyId"`
	BodyName string `json:"bodyName"`
	CampId   int    `json:"campId"`
	CampName string `json:"campName"`
	Data     struct {
		Dungeons map[string]map[string]struct {
			Pieces struct {
				Speed int `json:"speed"`
				Total int `json:"total"`
			} `json:"pieces"`
			Seniority struct {
				Speed int `json:"speed"`
				Total int `json:"total"`
			} `json:"seniority"`
		} `json:"dungeons"`
		Maps map[string]struct {
			Pieces struct {
				Speed int `json:"speed"`
				Total int `json:"total"`
			} `json:"pieces"`
			Seniority struct {
				Speed int `json:"speed"`
				Total int `json:"total"`
			} `json:"seniority"`
		} `json:"maps"`
		Score int `json:"score"`
		Total map[string]map[string]struct {
			Pieces struct {
				Speed int `json:"speed"`
				Total int `json:"total"`
			} `json:"pieces"`
			Seniority struct {
				Speed int `json:"speed"`
				Total int `json:"total"`
			} `json:"seniority"`
		} `json:"total"`
		TotalScore int `json:"totalScore"`
	} `json:"data"`
	ForceId    int    `json:"forceId"`
	ForceName  string `json:"forceName"`
	GlobalId   string `json:"globalId"`
	RoleId     string `json:"roleId"`
	RoleName   string `json:"roleName"`
	ServerName string `json:"serverName"`
	TongId     int    `json:"tongId"`
	TongName   string `json:"tongName"`
	ZoneName   string `json:"zoneName"`
}

// RoleAchievement 查询已在剑网3魔盒绑定角色的成就和资历分布。
func (c *Client) RoleAchievement(ctx context.Context, server string, name string, class int, subclass string, token string) (*RoleAchievementResponse, error) {
	params := &struct {
		Server   string `json:"server"`
		Name     string `json:"name"`
		Class    int    `json:"class"`
		Subclass string `json:"subclass,omitempty"`
		Token    string `json:"token,omitempty"`
	}{
		Server:   server,
		Name:     name,
		Class:    class,
		Subclass: subclass,
		Token:    token,
	}
	return requestData[RoleAchievementResponse](ctx, c, "/role/achievement", params)
}
