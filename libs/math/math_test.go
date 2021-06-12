package math

import (
	"testing"
	"wxl/core"
	"wxl/element"
	"wxl/language"
	"wxl/runtime"
)

func runCode(s string) element.Element {
	spec := language.CreateSpec([]language.wxlMethod{
		Add,
		Substract,
	})

	code := core.Parse(core.Tokenize(string(s)))

	return runtime.Run(spec, code)
}

func TestAdd(t *testing.T) {
	result := runCode("(+ 2 3)")

	if result.IsError() || result.NumberValue() != 5 {
		t.Error(result.StringValue())
	}
}
