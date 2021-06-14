package memory

import (
	"wxl/directives"
	"wxl/element"
	"wxl/language"
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

func (bind ElementBind) IsKeywordBind() bool {
	return false
}

func (bind ElementBind) GetElementValue() element.Element {
	return *bind.element
}

func (bind ElementBind) GetMethodValue() language.Method {
	el := *bind.element
	s := el.SymbolElementValue()

	return directives.Method{
		Symbol: *s,
		Call: func(params []*element.Element, ctx *language.Context) element.Element {
			return *params[len(params)-1]
		},
	}
}

func (bind ElementBind) GetKeywordValue() language.Keyword {
	return directives.Keyword{
		Call: func(params []*element.Element, ctx *language.Context) element.Element {
			return *params[len(params)-1]
		},
	}
}
