package libs

import (
	"fmt"
	"testing"
	"wxl/model"
)

func TestEntity(t *testing.T) {
	result := runCode(`
		(entity (model task label: string done: bool) label: "test123" done: true)
	`)

	if !result.IsObject() {
		t.Error("Result is: " + result.StringValue())

		return
	}

	ObjectElement := result.ObjectElementValue()

	fmt.Print(ObjectElement)

	object := *ObjectElement.Object()
	entity, ok := object.(model.Entity)

	if !ok {
		t.Error("Result is not entity.")
		return
	}

	if entity.Model.Name() != "task" {
		t.Error("Wront model name in the entity.")
		return
	}

	if entity.Attribute("label").Value.StringValue() != "test123" {
		t.Error("Wrong attribute value")
		return
	}

	if entity.Attribute("done").Value.BoolValue() == false {
		t.Error("Wrong attribute value")
		return
	}
}
