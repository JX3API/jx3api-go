package jx3api

import (
	"context"
	"testing"
)

func TestActiveMonster(t *testing.T) {
	client := NewClient(nil)

	client.ActiveMonster(context.TODO())
}

func TestHorseRecord(t *testing.T) {
	client := NewClient(nil)

	client.HorseRecord(context.TODO(), "唯我独尊")
}

func TestHorseRanch(t *testing.T) {
	client := NewClient(nil)

	client.HorseRanch(context.TODO(), "唯我独尊")
}

func TestFireworkRecord(t *testing.T) {
	client := NewClient(nil)

	client.FireworkRecord(context.TODO(), "唯我独尊", "风月的男宠")
}

func TestFireworkStatistical(t *testing.T) {
	client := NewClient(nil)

	client.FireworkStatistical(context.TODO(), "唯我独尊", "真橙之心", 5)
}

func TestFireworkRankStatistical(t *testing.T) {
	client := NewClient(nil)

	client.FireworkRankStatistical(context.TODO(), "唯我独尊", "sender", 1640285289, 1640687624)
}
