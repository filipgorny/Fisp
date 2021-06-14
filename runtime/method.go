package runtime

import "wxl/element"

type Method struct {
	Symbol element.SymbolElement
	Call   Call
}
