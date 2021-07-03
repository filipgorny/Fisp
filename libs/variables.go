package libs

import (
	"wxl/directives"
	"wxl/element"
	"wxl/language"
	"wxl/memory"
)

var Declare = directives.Method{
	Symbol: element.SymbolElement{Value: "set"},
	Call: func(params []*element.Element, ctx *language.Context) element.Element {
		currentContext := *ctx

		if len(params) < 3 {
			return element.ErrorElement{}
		}

		name := *params[1]
		value := params[2]

		currentContext.Declare(*name.SymbolElementValue(), memory.NewElementBind(value))

		return *value
	},
}
