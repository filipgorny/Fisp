package wxl

import (
	"io/ioutil"
	"strings"
)

const (
	LT_OPEN_LIST  = iota
	LT_CLOSE_LIST = iota
	LT_STRING     = iota
	LT_NUMBER     = iota
	LT_PATH       = iota
	LT_ATOM       = iota
	LT_ESCAPE     = iota
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("./lisp/test1.wxl")
	check(err)

	tokenize(string(dat))
}

func tokenize(code string) {
	code = strings.ReplaceAll(code, "`(", " (quote ")
	code = strings.ReplaceAll(code, "(", " ( ")
	code = strings.ReplaceAll(code, "(", " ( ")
	code = strings.Trim(code, " ")

	tokens := strings.Split(code, " ")

	for _, token := range tokens {

	}
}

func readToken(token string) {

}
