package libs

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	result := runCode("(+ 2 3)")

	if result.IsError() || result.NumberValue() != 5 {
		t.Error(result.StringValue())
	}
}

func TestDivMult(t *testing.T) {
	result := runCode("(/ (* 2 3) 3)")

	if result.IsError() || result.NumberValue() != 2 {
		t.Error(result.StringValue())
	}
}

func TestPerformance(t *testing.T) {
	code := ""

	for i := 1; i < 100; i++ {
		code += "(+ 2 3)"
	}

	code += "(- 5000 1000)"

	result := runCode(code)

	fmt.Print(result)

	if result.IsError() || result.NumberValue() == 490023 {
		t.Error(result.StringValue())
	}
}
