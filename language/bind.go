package language

import (
	"wxl/element"
	"wxl/object"
)

type Bind interface {
	IsElementBind() bool
	IsMethodBind() bool
	IsKeywordBind() bool
	IsObjectBind() bool
	GetElementValue() element.Element
	GetMethodValue() Method
	GetKeywordValue() Keyword
	GetObjectValue() *object.Object
}
