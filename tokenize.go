package main

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	LT_OPEN_LIST  = iota
	LT_CLOSE_LIST = iota
	LT_STRING     = iota
	LT_NUMBER     = iota
	LT_SYMBOL     = iota
)

type LToken struct {
	kind        int
	stringValue string
	numberValue float64
}

func tokenize(code string) []LToken {
	// code = strings.ReplaceAll(code, "`(", " (quote ")
	code = strings.ReplaceAll(code, "(", " ( ")
	code = strings.ReplaceAll(code, ")", " ) ")
	// code = strings.Trim(code, " ")
	tokens := strings.Split(code, " ")

	var result []LToken

	for _, token := range tokens {
		token = strings.Trim(token, " ")

		if token == "" {
			continue
		}

		result = append(result, readToken(token))
	}

	return result
}

func readToken(token string) LToken {
	var validNumber = regexp.MustCompile(`[+-]?[0-9]+(\\.[0-9]+)?([Ee][+-]?[0-9]+)?`)

	if token == "(" {
		return LToken{LT_OPEN_LIST, "(", 0}
	} else if token == ")" {
		return LToken{LT_CLOSE_LIST, ")", 0}
	} else if validNumber.MatchString(token) {
		number, _ := strconv.ParseFloat(token, 64)
		return LToken{LT_NUMBER, token, number}
	} else {
		return LToken{LT_SYMBOL, token, 0}
	}
}
