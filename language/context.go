package language

import "wxl/element"

type Context interface {
	Branch() Context
	Declare(s element.SymbolElement, bind Bind)
	Lookup(s element.Element) Bind
	Log(content interface{})
	Error(content interface{})
}
