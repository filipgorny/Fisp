package language

import "wxl/element"

type Keyword interface {
	GetSymbol() element.SymbolElement
	GetCall() Call
}
