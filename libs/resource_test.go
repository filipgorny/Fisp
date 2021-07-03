package libs

import (
	"testing"
)

func TestPut(t *testing.T) {
	result := runCode(`
		(put .task1 (entity (model task (field label: ^string) (field done: ^bool)) label: "test123" done: true))
    (get .task1)
	`)

	if !result.IsObject() {
		t.Error("Result is not the object.")

		return
	}

	if result.ObjectElementValue().Name.Value != "entity" {
		t.Error("Got object that is not an entity.")

		return
	}
}
