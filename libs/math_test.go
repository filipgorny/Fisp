package libs

import (
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
