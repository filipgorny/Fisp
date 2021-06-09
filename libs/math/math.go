package math

import (
	"wxl/core"
	"wxl/language"
)

var Add = language.BuiltInMethod{
	Symbol: core.SymbolElement{Value: "+"},
	Call: func(params []*core.Element) language.BuiltInMethodCallResult {
		var sum = 0.0

		var tokenElement core.Element

		for i, token := range params {
			tokenElement = *token

			if i == 0 {
				continue
			}

			sum += tokenElement.NumberValue()
		}

		return language.BuiltInMethodCallResult{Element: core.NumberElement{Value: sum}}
	},
}

var Substract = language.BuiltInMethod{
	Symbol: core.SymbolElement{Value: "-"},
	Call: func(params []*core.Element) language.BuiltInMethodCallResult {
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

		return language.BuiltInMethodCallResult{Element: core.NumberElement{Value: sum}}
	},
}
