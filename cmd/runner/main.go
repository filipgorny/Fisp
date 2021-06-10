package main

import (
	"io/ioutil"
	"rpl/core"
	"rpl/language"
	"rpl/libs/math"
	"rpl/runtime"
)

func main() {
	// arg := os.Args[1]
	arg := "lisp/test1.rpl"
	dat, err := ioutil.ReadFile(arg)

	if err != nil {
		panic(err)
	}

	spec := language.CreateSpec([]language.RplMethod{
		math.Add,
		math.Substract,
	})

	code := core.Parse(core.Tokenize(string(dat)))

	runtime.Run(spec, code)
}
