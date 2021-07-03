package language

import (
	"wxl/element"
	"wxl/resource"
)

type Context interface {
	resource.Resource

	Declare(s element.SymbolElement, bind Bind)
	Lookup(s element.Element) Bind

	Put(p element.PathElement, value resource.Resource)
	Get(p element.PathElement) resource.Resource

	Branch() *Context
}
