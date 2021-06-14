package math

import (
	"testing"
	"wxl/core"
	"wxl/element"
	"wxl/execution"
	"wxl/runtime"
)

func runCode(s string) element.Element {
	env := runtime.NewEnvironment([]*runtime.Method{
		&Add,
		&Substract,
	})

	code := core.Parse(core.Tokenize(string(s)))

	return execution.Run(env, code)
}

func TestAdd(t *testing.T) {
	result := runCode("(+ 2 3)")

	if result.IsError() || result.NumberValue() != 5 {
		t.Error(result.StringValue())
	}
}
