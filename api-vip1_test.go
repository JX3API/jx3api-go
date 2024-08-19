package jx3api

import (
	"context"
	"os"
	"testing"
)

func TestSaveDetailed(t *testing.T) {
	client := NewClient(nil)

	client.SaveDetailed(context.TODO(), "梦江南", "")
}

func TestRoleDetailed(t *testing.T) {
	client := NewClient(nil)

	client.RoleDetailed(context.TODO(), "梦江南", "")
}

func TestSchoolMatrix(t *testing.T) {
	client := NewClient(nil)

	client.SchoolMatrix(context.TODO(), "冰心诀")
}

func TestSchoolForce(t *testing.T) {
	client := NewClient(nil)

	client.SchoolForce(context.TODO(), "花间游")
}

func TestSchoolSkills(t *testing.T) {
	client := NewClient(nil)

	client.SchoolSkills(context.TODO(), "花间游")
}

func TestRoleTeamCdList(t *testing.T) {
	client := NewClient(nil)

	client.RoleTeamCdList(context.TODO(), "梦江南", "")
}

func TestLuckAdventure(t *testing.T) {
	client := NewClient(nil)

	client.LuckAdventure(context.TODO(), "梦江南", "狸嫁")
}

func TestLuckStatistical(t *testing.T) {
	client := NewClient(nil)

	client.LuckStatistical(context.TODO(), "长安城", "阴阳两界", 1)
}

func TestLuckServerStatistical(t *testing.T) {
	client := NewClient(nil)

	client.LuckServerStatistical(context.TODO(), "阴阳两界", 20)
}

func TestLuckCollect(t *testing.T) {
	client := NewClient(nil)

	client.LuckCollect(context.TODO(), "梦江南", 7)
}

func TestRoleAchievement(t *testing.T) {
	client := NewClient(nil)

	client.RoleAchievement(context.TODO(), "唯我独尊", "夜温言@长安城", "阴阳两界")
}

func TestMatchRecent(t *testing.T) {
	client := NewClient(nil)

	client.MatchRecent(context.TODO(), "梦江南", "有所依", 33)
}

func TestMatchAwesome(t *testing.T) {
	client := NewClient(nil)

	client.MatchAwesome(context.TODO(), 33, 2)
}

func TestMatchSchools(t *testing.T) {
	client := NewClient(nil)

	client.MatchSchools(context.TODO(), 33)
}

func TestMemberRecruit(t *testing.T) {
	client := NewClient(nil)

	client.MemberRecruit(context.TODO(), "梦江南", "英雄冷龙峰", 1)
}

func TestMemberTeacher(t *testing.T) {
	client := NewClient(nil)

	client.MemberTeacher(context.TODO(), "长安城", nil)
}

func TestMemberStudent(t *testing.T) {
	client := NewClient(nil)

	client.MemberStudent(context.TODO(), "长安城", nil)
}

func TestServerSand(t *testing.T) {
	client := NewClient(nil)

	client.ServerSand(context.TODO(), "长安城")
}

func TestServerEvent(t *testing.T) {
	client := NewClient(nil)

	client.ServerEvent(context.TODO(), "恶人谷", 100)
}

func TestDemonPrice(t *testing.T) {
	client := NewClient(nil)

	client.DemonPrice(context.TODO(), "长安城", 1)
}

func TestTradeRecord(t *testing.T) {
	client := NewClient(nil)

	client.TradeRecord(context.TODO(), "狐金")
}

func TestTiebaItemRecords(t *testing.T) {
	client := NewClient(nil)

	client.TiebaItemRecord(context.TODO(), "狐金", "长安城", 10)
}

func TestValuablesStatistical(t *testing.T) {
	client := NewClient(nil)

	client.ValuablesStatistical(context.TODO(), "长安城", "太一玄晶", 10)
}

func TestValuablesServerStatistical(t *testing.T) {
	client := NewClient(nil)

	client.ValuablesServerStatistical(context.TODO(), "太一玄晶", 15)
}

func TestRankStatistical(t *testing.T) {
	client := NewClient(nil)

	client.RankStatistical(context.TODO(), "长安城", "个人", "名士五十强")
}

func TestSchoolRankStatistical(t *testing.T) {
	client := NewClient(nil)

	client.SchoolRankStatistical(context.TODO(), "万花", "梦江南")
}
