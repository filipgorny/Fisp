package main

import (
	"io/ioutil"
	"wxl/core"
	"wxl/directives"
	"wxl/execution"
	"wxl/libs"
	"wxl/runtime"
)

func main() {
	// arg := os.Args[1]
	arg := "lisp/test1.wxl"
	dat, err := ioutil.ReadFile(arg)

	if err != nil {
		panic(err)
	}

	env := runtime.NewEnvironment([]*directives.Method{
		&libs.Add,
		&libs.Substract,
		&libs.Declare,
	}, []*directives.Keyword{
		&libs.Fun,
	})

	code := core.Parse(core.Tokenize(string(dat)))

	execution.Run(env, code)
}
