package jx3api

import (
	"context"
	"testing"
)

func TestActivateCalendar(t *testing.T) {
	client := NewClient(nil)

	client.ActivateCalendar(context.TODO(), "梦江南", 0)
}

func TestActivateListCalendar(t *testing.T) {
	client := NewClient(nil)

	client.ActivateListCalendar(context.TODO(), 7)
}

func TestActivateCelebrity(t *testing.T) {
	client := NewClient(nil)

	client.ActiveCelebs(context.TODO(), "楚天社")
}

func TestExamAnswer(t *testing.T) {
	client := NewClient(nil)

	client.ExamAnswer(context.TODO(), "李白", 10)
}

func TestHomeFurniture(t *testing.T) {
	client := NewClient(nil)

	client.HomeFurniture(context.TODO(), "龙门香梦")
}

func TestHomeTravel(t *testing.T) {
	client := NewClient(nil)

	client.HomeTravel(context.TODO(), "万花")
}

func TestNewsAllnews(t *testing.T) {
	client := NewClient(nil)

	client.NewsAllnews(context.TODO(), 1)
}

func TestNewsAnnounce(t *testing.T) {
	client := NewClient(nil)

	client.NewsAnnounce(context.TODO(), 1)
}

func TestServerMaster(t *testing.T) {
	client := NewClient(nil)

	client.ServerMaster(context.TODO(), "梦江南")
}

func TestStatusCheck(t *testing.T) {
	client := NewClient(nil)

	client.StatusCheck(context.TODO(), 1, "梦江南")
}

func TestHomeFlower(t *testing.T) {
	client := NewClient(nil)

	client.HomeFlower(context.TODO(), "梦江南", "", "")
}

func TestSkillRework(t *testing.T) {
	client := NewClient(nil)

	client.SkillRework(context.TODO())
}

func TestSchoolFoods(t *testing.T) {
	client := NewClient(nil)

	client.SchoolFoods(context.TODO())
}
