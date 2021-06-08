package main

import (
	"io/ioutil"
)

func main() {
	dat, err := ioutil.ReadFile("./lisp/test1.wxl")

	if err != nil {
		panic(err)
	}

	tokens := tokenize(string(dat))

	parse(tokens)
}
