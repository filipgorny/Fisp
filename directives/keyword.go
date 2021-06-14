package directives

import (
	"wxl/element"
	"wxl/language"
)

type Keyword struct {
	Symbol element.SymbolElement
	Call   language.Call
}

func (k Keyword) GetSymbol() element.SymbolElement {
	return k.Symbol
}

func (k Keyword) GetCall() language.Call {
	return k.Call
}
