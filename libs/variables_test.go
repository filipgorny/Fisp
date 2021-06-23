package libs

import (
	"testing"
)

func TestSet(t *testing.T) {
	result := runCode("(set x 10) (+ 20 x)")

	if result.IsError() || result.NumberValue() != 30 {
		t.Error(result.StringValue())
	}
}
