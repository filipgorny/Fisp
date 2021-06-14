package language

import "wxl/element"

type Method interface {
	GetSymbol() element.SymbolElement
	GetCall() Call
}
