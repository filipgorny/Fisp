package libs

import (
	"testing"
)

func TestFun(t *testing.T) {
	result := runCode("(fun test (a b) ((set x 30) (+ a b x))) (test 10 20)")

	if result.IsError() || result.NumberValue() != 60 {
		t.Error(result.StringValue())
	}
}

func TestFunAnonymous(t *testing.T) {
	result := runCode("((fun (a b) (+ a b 5)) 10 20)")

	if result.IsError() || result.NumberValue() != 35 {
		t.Error(result.StringValue())
	}
}

func TestFunComplex(t *testing.T) {
	result := runCode("(fun test (a b) ((set x 30) (+ a b x))) (fun subtest (a b) (- a b)) (subtest (test 10 20) 20)")

	if result.IsError() || result.NumberValue() != 40 {
		t.Error(result.StringValue())
	}
}
