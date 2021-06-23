package model

import (
	"wxl/element"
	"wxl/object"
)

type View struct {
	name element.StringElement
}

type Element struct {
	name       element.StringElement
	parameters []*Parameter
}

type Parameter struct {
	name  element.StringElement
	value element.Element
}

func (v View) Class() object.Class {
	return object.Class{Name: "view"}
}

func (e Element) Class() object.Class {
	return object.Class{Name: "view-element"}
}

func (p Parameter) Class() object.Class {
	return object.Class{Name: "view-parameter"}
}

func (v View) Name() element.StringElement {
	return v.name
}

func (e Element) Name() element.StringElement {
	return e.name
}

func (e Element) Parameters() []*Parameter {
	return e.parameters
}

func (p Parameter) Name() element.StringElement {
	return p.name
}

func (p Parameter) Value() element.Element {
	return p.value
}

func NewView(name string) View {
	return View{
		name: element.NewStringElement(name),
	}
}

func NewElement(name string) Element {
	return Element{
		name: element.NewStringElement(name),
	}
}

func (e *Element) AddParameter(p Parameter) {
	e.parameters = append(e.parameters, &p)
}

func NewParameter(name string, value element.Element) Parameter {
	return Parameter{
		name:  element.NewStringElement(name),
		value: value,
	}
}
