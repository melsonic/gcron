package util

import "fmt"

type Cron struct {
	Minute     string
	Hour       string
	DayOfMonth string
	Month      string
	DayOfWeek  string
}

func (c *Cron) Print() {
	fmt.Printf("%-15s : %s\n", "minute", c.Minute)
	fmt.Printf("%-15s : %s\n", "hour", c.Hour)
	fmt.Printf("%-15s : %s\n", "day_of_month", c.DayOfMonth)
	fmt.Printf("%-15s : %s\n", "month", c.Month)
	fmt.Printf("%-15s : %s\n", "day_of_week", c.DayOfWeek)
  fmt.Println()
}
