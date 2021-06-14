package runtime

import (
	"wxl/element"
	"wxl/language"
	"wxl/memory"
)

func Evaluate(listElement *element.ListElement, ctx language.Context) element.Element {
	firstElement := *listElement.First()

	ctx.Log("First element: " + firstElement.StringValue())

	if firstElement.IsSymbol() {
		bind := ctx.Lookup(firstElement.SymbolElementValue())

		if bind.IsMethodBind() {
			ctx.Log("method bind: " + bind.GetMethodValue().GetSymbol().Value)
			var params []*element.Element

			for _, param := range listElement.Children() {
				el := *param

				if el.IsList() {
					ctx.Log("Branching list: " + el.StringValue())
					newCtx := ctx.Branch()
					el = Evaluate(el.ListElementValue(), newCtx)
				}

				if el.IsSymbol() {
					elBind := ctx.Lookup(el.SymbolElementValue())

					if elBind.IsElementBind() {
						ctx.Log("Symbol bind: " + elBind.GetElementValue().StringValue())
						el = elBind.GetElementValue()
					}
				}

				params = append(params, &el)
			}

			return bind.GetMethodValue().GetCall()(params, &ctx)
		}

		if bind.IsKeywordBind() {
			ctx.Log("keyword bind: " + bind.GetMethodValue().GetSymbol().Value)
			var params []*element.Element

			for _, param := range listElement.Children() {
				el := *param

				// if el.IsSymbol() {
				// 	elBind := ctx.Lookup(el.SymbolElementValue())

				// 	if elBind.IsElementBind() {
				// 		ctx.Log("Symbol bind: " + elBind.GetElementValue().StringValue())
				// 		el = elBind.GetElementValue()
				// 	}
				// }

				params = append(params, &el)
			}

			return bind.GetKeywordValue().GetCall()(params, &ctx)
		}

		if bind.GetElementValue().IsFunction() {
			ctx.Log("function bind: " + bind.GetMethodValue().GetSymbol().Value)
			var params []*element.Element

			for _, param := range listElement.Children() {
				el := *param

				// if el.IsSymbol() {
				// 	elBind := ctx.Lookup(el.SymbolElementValue())

				// 	if elBind.IsElementBind() {
				// 		ctx.Log("Symbol bind: " + elBind.GetElementValue().StringValue())
				// 		el = elBind.GetElementValue()
				// 	}
				// }

				params = append(params, &el)
			}

			ctx.Log("Function exec")

			functionElement := bind.GetElementValue().FunctionElementValue()
			functionCtx := ctx.Branch()

			children := functionElement.GetArguments().Children()

			for i, _ := range children {

				el := *children[i]
				functionCtx.Declare(*el.SymbolElementValue(), memory.NewElementBind(params[i+1]))

				ctx.Log("bind for function : " + el.StringValue() + ", " + (*params[i+1]).StringValue())
			}

			resultOfFunction := Evaluate(functionElement.GetBody().ListElementValue(), functionCtx)

			return resultOfFunction
		}
	}

	last := *listElement.Last()

	for last.IsList() {
		result := element.ListElement{}

		for _, param := range listElement.Children() {
			el := *param

			if el.IsList() {
				ctx.Log("Branching list: " + el.StringValue())
				newCtx := ctx.Branch()
				el = Evaluate(el.ListElementValue(), newCtx)
			}

			if el.IsSymbol() {
				elBind := ctx.Lookup(el.SymbolElementValue())

				if elBind.IsElementBind() {
					ctx.Log("Symbol bind: " + elBind.GetElementValue().StringValue())
					el = elBind.GetElementValue()
				}
			}

			result.Push(el)
		}

		last = *result.Last()
	}

	return last
}
