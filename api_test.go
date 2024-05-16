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

	client.ActivateCelebrity(context.TODO(), 3)
}
