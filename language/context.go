package language

import (
	"wxl/element"
	"wxl/resource"
)

type Context interface {
	resource.Resource

	Branch() Context
	Declare(s element.SymbolElement, bind Bind)
	Lookup(s element.Element) Bind
	Log(content interface{})
	Error(content interface{})
}
