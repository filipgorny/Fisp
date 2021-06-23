package core

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
	LT_LABEL      = iota
	LT_PATH       = iota
)

type LToken struct {
	kind        int
	stringValue string
	numberValue float64
}

func Tokenize(code string) []LToken {
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
	var validString = regexp.MustCompile(`"(.*?)"`)
	var validLabel = regexp.MustCompile(`([a-z]+)\:`)
	var validPath = regexp.MustCompile(`\.([a-z]+)`)

	if token == "(" {
		return LToken{LT_OPEN_LIST, "(", 0}
	} else if token == ")" {
		return LToken{LT_CLOSE_LIST, ")", 0}
	} else if validString.MatchString(token) {
		stringValue := strings.Trim(token, `"`)
		return LToken{LT_STRING, stringValue, 0}
	} else if validNumber.MatchString(token) {
		number, _ := strconv.ParseFloat(token, 64)
		return LToken{LT_NUMBER, token, number}
	} else if validLabel.MatchString(token) {
		stringValue := strings.Trim(token, `:`)
		return LToken{LT_LABEL, stringValue, 0}
	} else if validPath.MatchString(token) {
		stringValue := strings.Trim(token, `:`)
		return LToken{LT_PATH, stringValue, 0}
	} else {
		return LToken{LT_SYMBOL, token, 0}
	}
}
