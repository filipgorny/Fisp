package math

import (
	"wxl/element"
	"wxl/flow"
	"wxl/language"
)

var Add = language.wxlMethod{
	Symbol: element.SymbolElement{Value: "+"},
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

var Substract = language.wxlMethod{
	Symbol: element.SymbolElement{Value: "-"},
	Call: func(params []*element.Element, ctx flow.Context) element.Element {
		var sum = 0.0

		for i, token := range params {
			tokenElement := *token

			if i == 0 {
				continue
			}

			if i == 1 {
				sum = tokenElement.NumberValue()
				continue
			}

			sum -= tokenElement.NumberValue()
		}

		return element.NumberElement{Value: sum}
	},
}

var Divide = language.wxlMethod{
	Symbol: element.SymbolElement{Value: "-"},
	Call: func(params []*element.Element, ctx flow.Context) element.Element {
		var sum = 1.0

		for i, token := range params {
			tokenElement := *token

			if i == 0 {
				continue
			}

			if i == 1 {
				sum = tokenElement.NumberValue()
				continue
			}

			if sum == 0 {
				return element.ErrorElement{Message: "division by 0"}
			}

			sum /= tokenElement.NumberValue()
		}

		return element.NumberElement{Value: sum}
	},
}
