package util

import (
	"testing"
)

func TestCronTabs(t *testing.T) {
	want := true
	c := Cron{
		Minute:     "1,3,5,21-49/3",
		Hour:       "0,3,5,7-22/3",
		DayOfMonth: "1-10",
		Month:      "AUG-DEC",
		DayOfWeek:  "MON-FRI",
	}
	got := c.Validate()
	if want != got {
		t.Errorf("got %t, want %t\n", got, want)
	}
}
