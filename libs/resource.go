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

		name := *params[1]
		resource := *params[2]

		currentContext.Put(*name.PathElementValue(), resource)

		return element.NullElement{}
	},
}

var Get = directives.Method{
	Symbol: element.SymbolElement{Value: "get"},
	Call: func(params []*element.Element, ctx *language.Context) element.Element {
		currentContext := *ctx

		name := *params[1]

		result := currentContext.Get(*name.PathElementValue())

		if result.(element.Element) != nil {
			return result.(element.Element)
		} else {
			return element.NullElement{}
		}
	},
}
