package language

import (
	"wxl/element"
)

type Bind interface {
	IsElementBind() bool
	IsMethodBind() bool
	IsKeywordBind() bool
	GetElementValue() element.Element
	GetMethodValue() Method
	GetKeywordValue() Keyword
}
