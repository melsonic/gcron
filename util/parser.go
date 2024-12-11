package util

import (
	"fmt"
	"strconv"
	"strings"
)

// main function to parse minutes
// precedance of operatore
// 1. comma
// 2. slash
// 3. dash
func ParseBase(input string, min_v int, max_v int) bool {
	var result bool = parseStar(input)
	if result {
		return true
	}
	result = parseComma(input, min_v, max_v)
	if result {
		return true
	}
	result = parseSlash(input, min_v, max_v)
	if result {
		return true
	}
	result = parseDash(input, min_v, max_v)
	if result {
		return true
	}
	result = parseValue(input, min_v, max_v)
	return result
}

// parse * i.e every minute
func parseStar(input string) bool {
	if input != "*" {
		return false
	}
	return true
}

// parse range of minutes
func parseDash(input string, min_v int, max_v int) bool {
	if !strings.Contains(input, "-") {
		return false
	}
	tokens := strings.Split(input, "-")
	// return if not in the form <left>-<right>
	if len(tokens) != 2 {
		return false
	}
	left, err := strconv.Atoi(tokens[0])
	right, err := strconv.Atoi(tokens[1])
	// return false in case of any of the following
	// components are not numbers
	// numbers are not in minute range
	if err != nil || left < min_v || left > max_v || right < min_v || right > max_v || left > right {
		return false
	}
	// if everything is fine store the parsed data into store
	return true
}

// parse step operator
func parseSlash(input string, min_v int, max_v int) bool {
	if !strings.Contains(input, "/") {
		return false
	}
	tokens := strings.Split(input, "/")
	// tokens must be of size 2
	// 1-49/5 will split into <1-49> and <5>
	if len(tokens) != 2 {
		return false
	}
	// first parse the second part
	_, error := strconv.Atoi(tokens[1])
	if error != nil {
		return false
	}
	// first try to parse dash if it has
	result := parseDash(tokens[0], min_v, max_v)
	// else it should be a single value or *
	// checking for *
	if !result {
		result = (tokens[0] == "*")
	}
	// checking for numerical minute value
	if !result {
		_, error = strconv.Atoi(tokens[0])
		if error == nil {
			result = true
		}
	}
	fmt.Println(result)
	return result
}

// parse comma operator
func parseComma(input string, min_v int, max_v int) bool {
	if !strings.Contains(input, ",") {
		return false
	}
	tokens := strings.Split(input, ",")
	var result bool
	for ix, token := range tokens {
		result = parseStar(token) || parseSlash(token, min_v, max_v) || parseDash(token, min_v, max_v) || parseValue(token, min_v, max_v)
		if ix != len(tokens)-1 {
		}
		if !result {
			break
		}
	}
	return result
}

func parseValue(input string, min_v int, max_v int) bool {
	num, error := strconv.Atoi(input)
	if error != nil || num < min_v || num > max_v {
		return false
	}
	return true
}

// parse Hour
// main thing is these are repetative things can we make them so that we can use them with other operators

// preprocessor functions
// preprocess month
// @input - user input crontab for months field which might contain JAN-DEC
func PreProcessMonth(input string) string {
	for month, id := range MONTHS_LIST {
		input = strings.ReplaceAll(input, month, strconv.Itoa(id))
	}
	return input
}

// preprocess week
// @input - user input crontab for week field which might contain SUN-MON
func PreProcessWeek(input string) string {
	for week, id := range WEEK_LIST {
		input = strings.ReplaceAll(input, week, strconv.Itoa(id))
	}
	// replace 7 with 0 if any
	input = strings.ReplaceAll(input, "7", "0")
	return input
}
