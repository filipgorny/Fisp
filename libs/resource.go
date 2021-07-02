package libs

import (
	"wxl/directives"
	"wxl/element"
	"wxl/language"
)

var Put = directives.Method{
	Symbol: element.SymbolElement{Value: "put"},
	Call: func(params []*element.Element, ctx *language.Context) element.Element {
		currentContext := *ctx

		return *value
	},
}
