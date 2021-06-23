package memory

import (
	"wxl/directives"
	"wxl/element"
	"wxl/language"
	"wxl/object"
)

type ObjectBind struct {
	object *object.Object
}

func NewObjectBind(object *object.Object) ObjectBind {
	return ObjectBind{
		object: object,
	}
}

func (bind ObjectBind) IsElementBind() bool {
	return false
}

func (bind ObjectBind) IsObjectBind() bool {
	return true
}

func (bind ObjectBind) IsMethodBind() bool {
	return false
}

func (bind ObjectBind) IsKeywordBind() bool {
	return false
}

func (bind ObjectBind) GetElementValue() element.Element {
	return element.NewObjectElement(*bind.object)
}

func (bind ObjectBind) GetMethodValue() language.Method {
	el := element.NewObjectElement(*bind.object)
	s := el.SymbolElementValue()

	return directives.Method{
		Symbol: *s,
		Call: func(params []*element.Element, ctx *language.Context) element.Element {
			return *params[len(params)-1]
		},
	}
}

func (bind ObjectBind) GetKeywordValue() language.Keyword {
	return directives.Keyword{
		Call: func(params []*element.Element, ctx *language.Context) element.Element {
			return *params[len(params)-1]
		},
	}
}

func (bind ObjectBind) GetObjectValue() *object.Object {
	return bind.object
}
