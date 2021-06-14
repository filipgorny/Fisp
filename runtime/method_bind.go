package runtime

import (
	"wxl/element"
)

type MethodBind struct {
	method Method
}

func (bind MethodBind) IsElementBind() bool {
	return false
}

func (bind MethodBind) IsMethodBind() bool {
	return true
}

func (bind MethodBind) GetElementValue() element.Element {
	return element.StringElement{Value: bind.method.Symbol.StringValue()}
}

func (bind MethodBind) GetMethodValue() Method {
	return bind.method
}
