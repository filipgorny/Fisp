package main

import (
	"io/ioutil"
	"wxl/core"
	"wxl/execution"
	"wxl/libs/math"
	"wxl/libs/variables"
	"wxl/runtime"
)

func main() {
	// arg := os.Args[1]
	arg := "lisp/test1.wxl"
	dat, err := ioutil.ReadFile(arg)

	if err != nil {
		panic(err)
	}

	env := runtime.NewEnvironment([]*runtime.Method{
		&math.Add,
		&math.Substract,
		&variables.Declare,
	})

	code := core.Parse(core.Tokenize(string(dat)))

	execution.Run(env, code)
}
