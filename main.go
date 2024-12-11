package main

import (
	"fmt"

	"github.com/melsonic/gcron/util"
)

func main() {
	util.PrintInfo()
	crontab := util.Cron{}
	// accept cron expression
	fmt.Printf("Cron Expression: ")
	fmt.Scan(&crontab.Minute, &crontab.Hour, &crontab.DayOfMonth, &crontab.Month, &crontab.DayOfWeek)
	if !crontab.Validate() {
		fmt.Println("Error parsing crontab")
		return
	}
	fmt.Println("parsing success!!!")
	crontab.Print()
}
