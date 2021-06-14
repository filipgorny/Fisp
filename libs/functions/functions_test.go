package functions

import (
	"testing"
	"wxl/core"
	"wxl/directives"
	"wxl/element"
	"wxl/execution"
	"wxl/libs/math"
	"wxl/libs/variables"
	"wxl/runtime"
)

func runCode(s string) element.Element {
	env := runtime.NewEnvironment([]*directives.Method{
		&math.Add,
		&variables.Declare,
	}, []*directives.Keyword{
		&Declare,
	})

	code := core.Parse(core.Tokenize(string(s)))

	return execution.Run(env, code)
}

func TestSet(t *testing.T) {
	result := runCode("(fun test (a b) ((set x 30) (+ a b x))) (test 10 20)")

	if result.IsError() || result.NumberValue() != 60 {
		t.Error(result.StringValue())
	}
}
