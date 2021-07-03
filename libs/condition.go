package libs

import (
	"wxl/directives"
	"wxl/element"
	"wxl/language"
)

var Equal = directives.Method{
	Symbol: element.SymbolElement{Value: "=="},
	Call: func(params []*element.Element, ctx *language.Context) element.Element {
		if len(params) < 2 {
			return element.NewBoolElement(true)
		}

		for i, token := range params {
			tokenElement := *token

			if i == 0 {
				continue
			}

			for ii, subToken := range params {
				subTokenElement := *subToken

				if ii < 2 {
					continue
				}

				if tokenElement != subTokenElement {
					return element.NewBoolElement(false)
				}
			}
		}

		return element.NewBoolElement(true)
	},
}

var NotEqual = directives.Method{
	Symbol: element.SymbolElement{Value: "!="},
	Call: func(params []*element.Element, ctx *language.Context) element.Element {

		if len(params) < 2 {
			return element.NewBoolElement(true)
		}

		for i, token := range params {
			tokenElement := *token

			if i == 0 {
				continue
			}

			for ii, subToken := range params {
				subTokenElement := *subToken

				if ii < 2 {
					continue
				}

				if tokenElement == subTokenElement {
					return element.NewBoolElement(false)
				}
			}
		}

		return element.NewBoolElement(true)
	},
}

var LeftBigger = directives.Method{
	Symbol: element.SymbolElement{Value: ">"},
	Call: func(params []*element.Element, ctx *language.Context) element.Element {
		if len(params) < 2 {
			return element.NewBoolElement(true)
		}

		for i, token := range params {
			tokenElement := *token

			if i == 0 {
				continue
			}

			for ii, subToken := range params {
				subTokenElement := *subToken

				if ii < 2 {
					continue
				}

				if tokenElement.NumberValue() < subTokenElement.NumberValue() && i < ii {
					return element.NewBoolElement(false)
				}
			}
		}

		return element.NewBoolElement(true)
	},
}

var RightBigger = directives.Method{
	Symbol: element.SymbolElement{Value: "<"},
	Call: func(params []*element.Element, ctx *language.Context) element.Element {
		if len(params) < 2 {
			return element.NewBoolElement(true)
		}

		for i, token := range params {
			tokenElement := *token

			if i == 0 {
				continue
			}

			for ii, subToken := range params {
				subTokenElement := *subToken

				if ii < 2 {
					continue
				}

				if tokenElement.NumberValue() > subTokenElement.NumberValue() && i < ii {
					return element.NewBoolElement(false)
				}
			}
		}

		return element.NewBoolElement(true)
	},
}

var IfTrue = directives.Method{
	Symbol: element.SymbolElement{Value: "?"},
	Call: func(params []*element.Element, ctx *language.Context) element.Element {
		if len(params) < 4 {
			return element.ErrorElement{Message: "Not enough elements."}
		}

		testEl := *params[1]

		do := *params[2]

		var elseDo element.Element

		elseDo = element.NullElement{}

		if len(params) > 3 {
			elseDo = *params[3]
		}

		if testEl.BoolValue() {
			return do
		} else {
			return elseDo
		}
	},
}
