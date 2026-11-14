package jx3api

import "context"

type MonsterRecordsResponse struct {
	GlobalId    string `json:"globalId"`
	RoleId      string `json:"roleId"`
	RoleName    string `json:"roleName"`
	Server      string `json:"server"`
	SkillCount  int    `json:"skillCount"`
	SkillEnergy int    `json:"skillEnergy"`
	SkillList   []struct {
		BDeprecated  bool   `json:"bDeprecated"`
		DwInSkillID  int    `json:"dwInSkillID"`
		DwOutSkillID int    `json:"dwOutSkillID"`
		NColor       int    `json:"nColor"`
		NCost        int    `json:"nCost"`
		NLevel       *int   `json:"nLevel,omitempty"`
		NSex         int    `json:"nSex"`
		SzBoss       string `json:"szBoss"`
		SzBossName   string `json:"szBossName"`
		SzSkillName  string `json:"szSkillName"`
		SzType       string `json:"szType"`
	} `json:"skillList"`
	SkillStamina int    `json:"skillStamina"`
	UpdateTime   int64  `json:"updateTime"`
	Zone         string `json:"zone"`
}

type RoleMonsterResponse = MonsterRecordsResponse

// MonsterRecords 查询角色的百战技能、体力、精力和场次。
func (c *Client) MonsterRecords(ctx context.Context, server string, name string, token string) (*MonsterRecordsResponse, error) {
	params := &struct {
		Server string `json:"server"`
		Name   string `json:"name"`
		Token  string `json:"token,omitempty"`
	}{
		Server: server,
		Name:   name,
		Token:  token,
	}
	return requestData[MonsterRecordsResponse](ctx, c, "/monster/records", params)
}

// RoleMonster is an alias for MonsterRecords.
func (c *Client) RoleMonster(ctx context.Context, server string, name string, token string) (*MonsterRecordsResponse, error) {
	return c.MonsterRecords(ctx, server, name, token)
}
