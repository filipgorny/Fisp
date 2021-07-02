package libs

import (
	"wxl/directives"
	"wxl/element"
	"wxl/language"
	"wxl/memory"
)

/*
(model model )


(new task (label: string))
*/

var Class = directives.Keyword{
	Symbol: element.SymbolElement{Value: "class"},
	Call: func(params []*element.Element, ctx *language.Context) element.Element {
		currentContext := *ctx
		firstArgument := *params[1]
		secondArgument := *params[2]

		if !firstArgument.IsSymbol() {
			return Error(ctx, "First argument must be a symbol.")
		}

		if len(params) < 3 {
			return Error(ctx, "Not enough arguments.")
		}

		object := *secondArgument.ObjectElementValue().Object()

		currentContext.Declare(*firstArgument.SymbolElementValue(), memory.NewObjectBind(&object))

		return secondArgument
	},
}
