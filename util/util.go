package util

import "fmt"

func PrintInfo() {
	fmt.Println("Please enter cron expression in the following format")
	fmt.Println("<minute> <hour> <day_of_month> <month> <day_of_week>")
	fmt.Println("where")
	fmt.Printf("%-15s : %s\n", "minute", "0-59")
	fmt.Printf("%-15s : %s\n", "hour", "0-23")
	fmt.Printf("%-15s : %s\n", "day_of_month", "1-31")
	fmt.Printf("%-15s : %s\n", "month", "1-12")
	fmt.Printf("%-15s : %s\n", "day_of_week", "0-7")
  fmt.Println()
}
