package math

import (
	"wxl/element"
	"wxl/runtime"
)

var Add = runtime.Method{
	Symbol: element.SymbolElement{Value: "+"},
	Call: func(params []*element.Element, ctx *runtime.Context) element.Element {
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

var Substract = runtime.Method{
	Symbol: element.SymbolElement{Value: "-"},
	Call: func(params []*element.Element, ctx *runtime.Context) element.Element {
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

var Divide = runtime.Method{
	Symbol: element.SymbolElement{Value: "-"},
	Call: func(params []*element.Element, ctx *runtime.Context) element.Element {
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
