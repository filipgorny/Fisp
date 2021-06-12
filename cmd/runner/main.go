package main

import (
	"io/ioutil"
	"wxl/core"
	"wxl/language"
	"wxl/libs/math"
	"wxl/runtime"
)

func main() {
	// arg := os.Args[1]
	arg := "lisp/test1.wxl"
	dat, err := ioutil.ReadFile(arg)

	if err != nil {
		panic(err)
	}

	spec := language.CreateSpec([]language.wxlMethod{
		math.Add,
		math.Substract,
	})

	code := core.Parse(core.Tokenize(string(dat)))

	runtime.Run(spec, code)
}
