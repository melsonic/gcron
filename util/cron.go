package util

import (
	"fmt"
)

type Cron struct {
	Minute     string
	Hour       string
	DayOfMonth string
	Month      string
	DayOfWeek  string
	Result     CronResult
}

type CronResult struct {
	Minute     string
	Hour       string
	DayOfMonth string
	Month      string
	DayOfWeek  string
}

// main crontab validator structure
func (c *Cron) Validate() bool {
	var result bool = true
	result = ParseBase(c.Minute, MIN_MINUTE, MAX_MINUTE)
	if !result {
		return false
	}
	result = ParseBase(c.Hour, MIN_HOUR, MAX_HOUR)
	if !result {
		return false
	}
	result = ParseBase(c.DayOfMonth, MIN_DAY_IN_MONTH, MAX_DAY_IN_MONTH)
	if !result {
		return false
	}
	// parse to change e.g JAN->1, FEB->2, MAR->3, ...
	c.Month = PreProcessMonth(c.Month)
	result = ParseBase(c.Month, MIN_MONTH, MAX_MONTH)
	if !result {
		return false
	}
	// parse to change e.g SUN->0, MON->1, TUE->3, ...
	c.DayOfWeek = PreProcessWeek(c.DayOfWeek)
	result = ParseBase(c.DayOfWeek, MIN_DAY_IN_WEEK, MAX_DAY_IN_WEEK)
	if !result {
		return false
	}
	return true
}

// print the crontab structore
func (c *Cron) Print() {
	fmt.Printf("%-15s : %s\n", "minute", c.Minute)
	fmt.Printf("%-15s : %s\n", "hour", c.Hour)
	fmt.Printf("%-15s : %s\n", "day_of_month", c.DayOfMonth)
	fmt.Printf("%-15s : %s\n", "month", c.Month)
	fmt.Printf("%-15s : %s\n", "day_of_week", c.DayOfWeek)
	fmt.Println()
}
