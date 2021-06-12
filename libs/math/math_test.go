package math

import (
	"rpl/core"
	"rpl/element"
	"rpl/language"
	"rpl/runtime"
	"testing"
)

func runCode(s string) element.Element {
	spec := language.CreateSpec([]language.RplMethod{
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
