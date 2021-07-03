package libs

import (
	"wxl/directives"
	"wxl/element"
	"wxl/language"
	"wxl/memory"
)

var Fun = directives.Keyword{
	Symbol: element.SymbolElement{Value: "fun"},
	Call: func(params []*element.Element, ctx *language.Context) element.Element {
		currentContext := *ctx

		if len(params) < 3 {
			return element.ErrorElement{}
		}

		var functionElement element.Element

		first := *params[1]
		second := *params[2]

		var third element.Element
		third = element.NullElement{}

		if len(params) > 3 {
			third = *params[3]
		}

		if first.IsSymbol() && second.IsList() && third.IsList() {
			functionElement = element.NewFunctionElement(
				true,
				first.StringValue(),
				*second.ListElementValue(),
				*third.ListElementValue(),
			)

			currentContext.Declare(*first.SymbolElementValue(), memory.NewElementBind(&functionElement))
		} else if first.IsList() && second.IsList() {
			functionElement = element.NewFunctionElement(
				false,
				first.StringValue(),
				*first.ListElementValue(),
				*second.ListElementValue(),
			)
		} else {
			msg := "Invalid function declaration."
			return element.ErrorElement{Message: msg}
		}

		return functionElement
	},
}
