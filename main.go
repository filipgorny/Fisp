package main

import (
	"io/ioutil"

	"github.com/kr/pretty"
)

func main() {
	dat, err := ioutil.ReadFile("./lisp/test1.wxl")

	if err != nil {
		panic(err)
	}

	tokens := tokenize(string(dat))

	ast := parse(tokens)

	pretty.Print(ast)
}
