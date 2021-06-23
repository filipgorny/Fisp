package core

import (
	"testing"
)

func runCode(s string) ExprList {
	return Parse(Tokenize(string(s)))
}

func TestRecord(t *testing.T) {
	result := runCode("(test test-record: test-record-value)").Elements[0]

	if result.IsError() || result.IsList() == false {
		t.Error(result.StringValue())
	}

	if !result.ListElementValue().Elements[1].IsRecord() {
		t.Error("Second argument is not record.")
	}

	if result.ListElementValue().Elements[1].RecordElementValue().Label() != "test-record" {
		t.Error("Wrong record label.")
	}

	if result.ListElementValue().Elements[1].RecordElementValue().Value().StringValue() != "test-record-value" {
		t.Error("Wrong record value.")
	}
}
