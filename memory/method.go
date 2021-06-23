package memory

import (
	"wxl/directives"
	"wxl/element"
	"wxl/language"
	"wxl/object"
)

type MethodBind struct {
	method language.Method
}

func NewMethodBind(method language.Method) MethodBind {
	return MethodBind{
		method: method,
	}
}

func (bind MethodBind) IsElementBind() bool {
	return false
}

func (bind MethodBind) IsObjectBind() bool {
	return false
}

func (bind MethodBind) IsMethodBind() bool {
	return true
}

func (bind MethodBind) IsKeywordBind() bool {
	return false
}

func (bind MethodBind) GetElementValue() element.Element {
	return element.StringElement{Value: bind.method.GetSymbol().StringValue()}
}

func (bind MethodBind) GetMethodValue() language.Method {
	return bind.method
}

func (bind MethodBind) GetKeywordValue() language.Keyword {
	return directives.Keyword{
		Call: func(params []*element.Element, ctx *language.Context) element.Element {
			return element.NullElement{}
		},
	}
}

func (bind MethodBind) GetObjectValue() *object.Object {
	return nil
}
