package runtime

import (
	"wxl/element"
)

type ElementBind struct {
	element *element.Element
}

func NewElementBind(element *element.Element) ElementBind {
	return ElementBind{
		element: element,
	}
}

func (bind ElementBind) IsElementBind() bool {
	return true
}

func (bind ElementBind) IsMethodBind() bool {
	return false
}

func (bind ElementBind) GetElementValue() element.Element {
	return *bind.element
}

func (bind ElementBind) GetMethodValue() Method {
	el := *bind.element
	s := el.SymbolElementValue()

	return Method{
		Symbol: *s,
		Call: func(params []*element.Element, ctx *Context) element.Element {
			return *params[len(params)-1]
		},
	}
}
