package directives

import (
	"wxl/element"
	"wxl/language"
)

type Method struct {
	Symbol element.SymbolElement
	Call   language.Call
}

func (m Method) GetSymbol() element.SymbolElement {
	return m.Symbol
}

func (m Method) GetCall() language.Call {
	return m.Call
}
