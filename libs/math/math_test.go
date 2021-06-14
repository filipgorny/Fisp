package math

import (
	"testing"
	"wxl/core"
	"wxl/directives"
	"wxl/element"
	"wxl/execution"
	"wxl/runtime"
)

func runCode(s string) element.Element {
	env := runtime.NewEnvironment([]*directives.Method{
		&Add,
		&Substract,
		&Divide,
		&Multiple,
	}, []*directives.Keyword{})

	code := core.Parse(core.Tokenize(string(s)))

	return execution.Run(env, code)
}

func TestAdd(t *testing.T) {
	result := runCode("(+ 2 3)")

	if result.IsError() || result.NumberValue() != 25 {
		t.Error(result.StringValue())
	}
}

func TestDivMult(t *testing.T) {
	result := runCode("(/ (* 2 3) 3)")

	if result.IsError() || result.NumberValue() != 2 {
		t.Error(result.StringValue())
	}
}
