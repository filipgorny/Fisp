package variables

import (
	"wxl/element"
	"wxl/runtime"
)

var Declare = runtime.Method{
	Symbol: element.SymbolElement{Value: "set"},
	Call: func(params []*element.Element, ctx *runtime.Context) element.Element {
		if len(params) < 3 {
			ctx.Error("Not enough arguments for method `set`.")

			return element.ErrorElement{}
		}

		name := *params[1]
		value := params[2]

		ctx.Log("Assigning " + name.StringValue() + " to " + (*value).StringValue())

		ctx.Declare(*name.SymbolElementValue(), runtime.NewElementBind(value))

		return *value
	},
}
