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

func (c *Cron) Validate() bool {
  var result bool = true
  result = ParseMinutesTab(c.Minute, &c.Result.Minute)
  fmt.Println(c.Result.Minute)
  if !result {
    return false
  }
  result = ParseHoursTab(c.Hour, &c.Result.Hour)
  return true
}

func (c *Cron) Print() {
	fmt.Printf("%-15s : %s\n", "minute", c.Minute)
	fmt.Printf("%-15s : %s\n", "hour", c.Hour)
	fmt.Printf("%-15s : %s\n", "day_of_month", c.DayOfMonth)
	fmt.Printf("%-15s : %s\n", "month", c.Month)
	fmt.Printf("%-15s : %s\n", "day_of_week", c.DayOfWeek)
	fmt.Printf("%-15s : %s\n", "minute result", c.Result.Minute)
	fmt.Println()
}
