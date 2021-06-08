package main

import (
	"fmt"
	"io/ioutil"
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

type Ast []LToken

var validNumber = regexp.MustCompile(`[+-]?[0-9]+(\\.[0-9]+)?([Ee][+-]?[0-9]+)?`)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("./lisp/test1.wxl")
	check(err)

	ast := tokenize(string(dat))

	fmt.Print(ast)
}

func tokenize(code string) Ast {
	// code = strings.ReplaceAll(code, "`(", " (quote ")
	code = strings.ReplaceAll(code, "(", " ( ")
	code = strings.ReplaceAll(code, ")", " ) ")
	code = strings.Trim(code, " ")

	tokens := strings.Split(code, " ")

	var ast Ast

	for _, token := range tokens {
		ast = append(ast, readToken((token)))
	}

	return ast
}

func readToken(token string) LToken {
	if token == " ( " {
		return LToken{LT_OPEN_LIST, "(", 0}
	} else if token == " ) " {
		return LToken{LT_CLOSE_LIST, ")", 0}
	} else if validNumber.MatchString(token) {
		number, _ := strconv.ParseFloat(token, 64)
		return LToken{LT_NUMBER, token, number}
	} else {
		return LToken{LT_SYMBOL, token, 0}
	}
}
