package util

import "fmt"

// min & max values of different crontab fields
const (
	MIN_MINUTE       = 0
	MAX_MINUTE       = 59
	MIN_HOUR         = 0
	MAX_HOUR         = 23
	MIN_DAY_IN_MONTH = 1
	MAX_DAY_IN_MONTH = 31
	MIN_MONTH        = 1
	MAX_MONTH        = 12
	MIN_DAY_IN_WEEK  = 0
	MAX_DAY_IN_WEEK  = 6
)

var (
	MONTHS_LIST = map[string]int{
		"JAN": 1,
		"FEB": 2,
		"MAR": 3,
		"APR": 4,
		"MAY": 5,
		"JUN": 6,
		"JUL": 7,
		"AUG": 8,
		"SEP": 9,
		"OCT": 10,
		"NOV": 11,
		"DEC": 12,
	}

	WEEK_LIST = map[string]int{
		"SUN": 0,
		"MON": 1,
		"TUE": 2,
		"WED": 3,
		"THU": 4,
		"FRI": 5,
		"SAT": 6,
	}
)

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
