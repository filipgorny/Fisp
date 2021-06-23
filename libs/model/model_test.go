package model

import (
	"fmt"
	"testing"
	"wxl/core"
	"wxl/directives"
	"wxl/element"
	"wxl/execution"
	"wxl/model"
	"wxl/runtime"
)

func runCode(s string) element.Element {
	env := runtime.NewEnvironment([]*directives.Method{
		&Model,
		&Entity,
	}, []*directives.Keyword{})

	code := core.Parse(core.Tokenize(string(s)))

	return execution.Run(env, code)
}

func TestEntity(t *testing.T) {
	result := runCode(`
		(entity (model task label: string done: bool) label: "test123" done: true)
	`)

	if !result.IsObject() {
		t.Error("Result is: " + result.StringValue())

		return
	}

	objectElement := result.ObjectElementValue()

	fmt.Print(objectElement)

	object := *objectElement.Object()
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
