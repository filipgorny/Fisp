package runtime

import (
	"wxl/element"
)

type Bind interface {
	IsElementBind() bool
	IsMethodBind() bool
	GetElementValue() element.Element
	GetMethodValue() Method
}

type NullBind struct {
}

func (bind NullBind) IsElementBind() bool {
	return false
}

func (bind NullBind) IsMethodBind() bool {
	return false
}

func (bind NullBind) GetElementValue() element.Element {
	return element.NullElement{}
}

func (bind NullBind) GetMethodValue() Method {
	return Method{
		Call: func(params []*element.Element, ctx *Context) element.Element {
			return element.NullElement{}
		},
	}
}
