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
func ParseMinutesTab(input string, store *string) bool {
	var result bool = parseStar(input, store)
	if result {
		return true
	}
	*store = ""
	result = parseComma(input, store)
	if result {
		return true
	}
	*store = ""
	result = parseSlash(input, store)
	if result {
		return true
	}
	*store = ""
	result = parseDash(input, store)
	if result {
		return true
	}
	*store = ""
	result = parseValue(input, store)
	return result
}

// parse * i.e every minute
func parseStar(input string, store *string) bool {
	if input != "*" {
		return false
	}
	*store += "every minute"
	return true
}

// parse numeric minute value
func parseValue(input string, store *string) bool {
	num, error := strconv.Atoi(input)
	if error != nil || num < 0 || num > 59 {
		return false
	}
	*store += fmt.Sprintf("at %dth minute", num)
	return true
}

// parse range of minutes
func parseDash(input string, store *string) bool {
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
	if err != nil || left < 0 || left > 59 || right < 0 || right > 59 || left > right {
		return false
	}
	// if everything is fine store the parsed data into store
	*store += fmt.Sprintf(" between %s and %s minutes ", tokens[0], tokens[1])
	return true
}

// parse step operator
func parseSlash(input string, store *string) bool {
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
	num, error := strconv.Atoi(tokens[1])
	if error != nil {
		return false
	}
	*store += fmt.Sprintf(" every %d minutes ", num)
	// first try to parse dash if it has
	result := parseDash(tokens[0], store)
	// else it should be a single value or *
  // checking for *
  if !result {
    result = (tokens[0] == "*")
  }
  // checking for numerical minute value
	if !result {
		num, error = strconv.Atoi(tokens[0])
		if error == nil {
			result = true
			*store += fmt.Sprintf(", starting at %d minutes past the hour ", num)
		}
	}
  fmt.Println(result)
  fmt.Println(*store)
	return result
}

// parse comma operator
func parseComma(input string, store *string) bool {
	if !strings.Contains(input, ",") {
		return false
	}
	tokens := strings.Split(input, ",")
	var result bool
	for ix, token := range tokens {
		result = parseStar(token, store) || parseSlash(token, store) || parseDash(token, store) || parseValue(token, store)
		if ix != len(tokens)-1 {
			*store += ", "
		}
		if !result {
			break
		}
	}
	return result
}

/// parse Hour
/// main thing is these are repetative things can we make them so that we can use them with other operators
