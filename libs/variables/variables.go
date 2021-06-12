package variables

import (
	"rpl/element"
	"rpl/flow"
	"rpl/language"
)

var Declare = language.RplMethod{
	Symbol: element.SymbolElement{Value: "set"},
	Call: func(params []*element.Element, ctx flow.Context) element.Element {
		var sum = 0.0

		var tokenElement element.Element

		for i, token := range params {
			tokenElement = *token

			if i == 0 {
				continue
			}

			sum += tokenElement.NumberValue()
		}

		return element.NumberElement{Value: sum}
	},
}
