package libs

import (
	"testing"
)

func TestEqual(t *testing.T) {
	result := runCode("(== 2 2)")

	if result.IsError() || result.NumberValue() != 1 {
		t.Error(result.StringValue())
	}
}

func TestEqualStrWithNum(t *testing.T) {
	result := runCode("(== 2 \"2\")")

	if result.IsError() || result.NumberValue() != 0 { // not yet implemented
		t.Error(result.StringValue())
	}
}

func TestLeftBigger(t *testing.T) {
	result := runCode("(> 5 2 1)")

	if result.IsError() || result.NumberValue() != 1 {
		t.Error(result.StringValue())
	}
}

func TestRightBigger(t *testing.T) {
	result := runCode("(< 1 2 3 4)")

	if result.IsError() || result.NumberValue() != 1 {
		t.Error(result.StringValue())
	}
}

func TestIfTrue(t *testing.T) {
	result := runCode("(? true 31 42)")

	if result.IsError() || result.NumberValue() != 31 {
		t.Error(result.StringValue())
	}
}

func TestIfTrueNot(t *testing.T) {
	result := runCode("(? (== 3 4) 31 42)")

	if result.IsError() || result.NumberValue() != 42 {
		t.Error(result.StringValue())
	}
}
