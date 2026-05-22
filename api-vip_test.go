package jx3api

import (
	"context"
	"testing"
)

// VIP 接口测试
func TestActiveMonster(t *testing.T) {
	client := NewClient(nil)
	client.ActiveMonster(context.TODO(), "")
}

func TestNextEvent(t *testing.T) {
	client := NewClient(nil)
	client.NextEvent(context.TODO(), "梦江南", "")
}

func TestAuctionRecord(t *testing.T) {
	client := NewClient(nil)
	client.AuctionRecord(context.TODO(), "梦江南", "物品名", 10, "")
}

func TestSteedRecord(t *testing.T) {
	client := NewClient(nil)
	client.SteedRecord(context.TODO(), "梦江南", "")
}

func TestShowRecord(t *testing.T) {
	client := NewClient(nil)
	client.ShowRecord(context.TODO(), "梦江南", "角色名", "")
}

func TestFraudDetail(t *testing.T) {
	client := NewClient(nil)
	client.FraudDetail(context.TODO(), 570790267, "")
}

func TestEventRecords(t *testing.T) {
	client := NewClient(nil)
	client.EventRecords(context.TODO(), "梦江南", "角色名", "")
}

func TestEventUnfinished(t *testing.T) {
	client := NewClient(nil)
	client.EventUnfinished(context.TODO(), "梦江南", "角色名", "")
}

func TestEventRecent(t *testing.T) {
	client := NewClient(nil)
	client.EventRecent(context.TODO(), "梦江南", "")
}

func TestEventStatistics(t *testing.T) {
	client := NewClient(nil)
	client.EventStatistics(context.TODO(), "梦江南", "角色名", 10, "")
}

func TestEventCollect(t *testing.T) {
	client := NewClient(nil)
	client.EventCollect(context.TODO(), "梦江南", 7, "")
}

func TestArenaRecent(t *testing.T) {
	client := NewClient(nil)
	client.ArenaRecent(context.TODO(), "梦江南", "角色名", 33, "", "")
}

func TestArenaAwesome(t *testing.T) {
	client := NewClient(nil)
	client.ArenaAwesome(context.TODO(), 33, 2, "", "")
}

func TestArenaSchools(t *testing.T) {
	client := NewClient(nil)
	client.ArenaSchools(context.TODO(), 33, "", "")
}

func TestRecruitSearch(t *testing.T) {
	client := NewClient(nil)
	client.RecruitSearch(context.TODO(), "梦江南", "关键词", 1, "")
}

func TestMentorSearch(t *testing.T) {
	client := NewClient(nil)
	client.MentorSearch(context.TODO(), 1, "长安城", "关键词", "")
}

func TestRankStatistical(t *testing.T) {
	client := NewClient(nil)
	client.RankStatistical(context.TODO(), "长安城", "名士五十强", "")
}

func TestRewardStatistics(t *testing.T) {
	client := NewClient(nil)
	client.RewardStatistics(context.TODO(), "梦江南", "物品名", 10, "")
}

func TestRoleDetail(t *testing.T) {
	client := NewClient(nil)
	client.RoleDetail(context.TODO(), "梦江南", "角色名", "")
}

func TestCardRecord(t *testing.T) {
	client := NewClient(nil)
	client.CardRecord(context.TODO(), "梦江南", "角色名", "")
}

func TestCardRecords(t *testing.T) {
	client := NewClient(nil)
	client.CardRecords(context.TODO(), "梦江南", "角色名", "")
}

func TestCardRandom(t *testing.T) {
	client := NewClient(nil)
	client.CardRandom(context.TODO(), "梦江南", "男性", "万花", "")
}

func TestCardCached(t *testing.T) {
	client := NewClient(nil)
	client.CardCached(context.TODO(), "梦江南", "角色名", "")
}

func TestRoleMonster(t *testing.T) {
	client := NewClient(nil)
	client.RoleMonster(context.TODO(), "梦江南", "角色名", "")
}

func TestSchoolMatrix(t *testing.T) {
	client := NewClient(nil)
	client.SchoolMatrix(context.TODO(), "冰心诀", "", "")
}

func TestSchoolTalent(t *testing.T) {
	client := NewClient(nil)
	client.SchoolTalent(context.TODO(), "冰心诀", "", "")
}

func TestSchoolSkills(t *testing.T) {
	client := NewClient(nil)
	client.SchoolSkills(context.TODO(), "花间游", "", "")
}

func TestSchoolSeniority(t *testing.T) {
	client := NewClient(nil)
	client.SchoolSeniority(context.TODO(), "梦江南", "万花", "", "")
}

func TestSandRecords(t *testing.T) {
	client := NewClient(nil)
	client.SandRecords(context.TODO(), "梦江南")
}

func TestFenxianRecords(t *testing.T) {
	client := NewClient(nil)
	client.FenxianRecords(context.TODO())
}

func TestSmiteRecords(t *testing.T) {
	client := NewClient(nil)
	client.SmiteRecords(context.TODO(), "")
}

func TestMineCart(t *testing.T) {
	client := NewClient(nil)
	client.MineCart(context.TODO(), "")
}

func TestChituRecords(t *testing.T) {
	client := NewClient(nil)
	client.ChituRecords(context.TODO(), "")
}

func TestChituWeekRecords(t *testing.T) {
	client := NewClient(nil)
	client.ChituWeekRecords(context.TODO(), "")
}

func TestRanchRecords(t *testing.T) {
	client := NewClient(nil)
	client.RanchRecords(context.TODO(), "梦江南", "")
}

func TestRankTrials(t *testing.T) {
	client := NewClient(nil)
	client.RankTrials(context.TODO(), "梦江南", "万花", "")
}

func TestTiebaItemRecords(t *testing.T) {
	client := NewClient(nil)
	client.TiebaItemRecords(context.TODO(), "梦江南", "狐金", 10, "")
}

func TestTradeDemon(t *testing.T) {
	client := NewClient(nil)
	client.TradeDemon(context.TODO(), "梦江南", 10, "")
}

func TestTradeRecords(t *testing.T) {
	client := NewClient(nil)
	client.TradeRecords(context.TODO(), "梦江南", "物品名", "")
}

func TestTradeItemSearch(t *testing.T) {
	client := NewClient(nil)
	client.TradeItemSearch(context.TODO(), "物品名", "")
}

func TestTradeItemRecords(t *testing.T) {
	client := NewClient(nil)
	client.TradeItemRecords(context.TODO(), "梦江南", "物品名", "")
}

func TestBattleRecords(t *testing.T) {
	client := NewClient(nil)
	client.BattleRecords(context.TODO(), "梦江南", "")
}

func TestMechCalculator(t *testing.T) {
	client := NewClient(nil)
	client.MechCalculator(context.TODO())
}

func TestDuowanStatistics(t *testing.T) {
	client := NewClient(nil)
	client.DuowanStatistics(context.TODO(), "梦江南", "")
}

// 其他接口测试
func TestTiebaRandom(t *testing.T) {
	client := NewClient(nil)
	client.TiebaRandom(context.TODO(), "818", "-", 10, "")
}

func TestSaohuaRandom(t *testing.T) {
	client := NewClient(nil)
	client.SaohuaRandom(context.TODO())
}

func TestSaohuaContent(t *testing.T) {
	client := NewClient(nil)
	client.SaohuaContent(context.TODO())
}

func TestSoundConverter(t *testing.T) {
	client := NewClient(nil)
	client.SoundConverter(context.TODO(), "appkey", "access", "secret", "Aitong", "mp3", 16000, 50, 0, 0, "测试文本")
}
