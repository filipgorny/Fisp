package variables

import (
	"testing"
	"wxl/core"
	"wxl/element"
	"wxl/execution"
	"wxl/libs/math"
	"wxl/runtime"
)

func runCode(s string) element.Element {
	env := runtime.NewEnvironment([]*runtime.Method{
		&Declare,
		&math.Add,
	})

	code := core.Parse(core.Tokenize(string(s)))

	return execution.Run(env, code)
}

func TestSet(t *testing.T) {
	result := runCode("(set x 10) (+ 20 x)")

	if result.IsError() || result.NumberValue() != 30 {
		t.Error(result.StringValue())
	}
}
