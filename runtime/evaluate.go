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

				params = append(params, &el)
			}

			bindResult := bind.GetKeywordValue().GetCall()(params, &ctx)

			ctx.Log("Keyword exec result: " + bindResult.StringValue())

			return bindResult
		}

		if bind.GetElementValue().IsFunction() {
			ctx.Log("function bind: " + bind.GetMethodValue().GetSymbol().Value)

			return EvaluateFun(*bind.GetElementValue().FunctionElementValue(), listElement, ctx)
		}
	}

	// anonymous function call

	if firstElement.IsFunction() {
		ctx.Log("Running anonymouse fun: " + firstElement.StringValue())
		newCtx := ctx.Branch()
		return EvaluateFun(*firstElement.FunctionElementValue(), listElement, newCtx)
	}

	// not callable list

	var last element.Element

	last = *listElement

	ctx.Log("Last el: " + last.StringValue())

	for last.IsList() {
		ctx.Log("Last el is a list, looping: " + last.StringValue())

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

	ctx.Log("Last is not a list, returing " + last.StringValue())

	return last
}

func EvaluateFun(functionElement element.FunctionElement, listElement *element.ListElement, ctx language.Context) element.Element {
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

	ctx.Log("Function exec")

	functionCtx := ctx.Branch()

	children := functionElement.GetArguments().Children()

	for i, _ := range children {

		el := *children[i]
		functionCtx.Declare(*el.SymbolElementValue(), memory.NewElementBind(params[i+1]))

		ctx.Log("bind for function : " + el.StringValue() + ", " + (*params[i+1]).StringValue())
	}

	return Evaluate(functionElement.GetBody().ListElementValue(), functionCtx)
}
