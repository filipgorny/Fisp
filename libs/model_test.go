package libs

import (
	"testing"
)

func TestEntity(t *testing.T) {
	result := runCode(`
		(entity (model task (field label: ^string) (field done: ^bool)) label: "test123" done: true)
	`)

	if !result.IsObject() {
		t.Error("Result is: " + result.StringValue())

		return
	}

	ObjectElement := result.ObjectElementValue()

	codeValue := ObjectElement.Serialize().StringValue()

	t.Log(codeValue)

	return
}
