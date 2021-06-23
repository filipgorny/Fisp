package runtime

import (
	"wxl/element"
	"wxl/language"
	"wxl/memory"
)

func Evaluate(listElement *element.ListElement, ctx language.Context) element.Element {
	firstElement := *listElement.First()

	if firstElement.IsSymbol() {
		bind := ctx.Lookup(firstElement.SymbolElementValue())

		if bind.IsMethodBind() {
			var params []*element.Element

			for _, param := range listElement.Children() {
				el := *param

				if el.IsList() {
					newCtx := ctx.Branch()
					el = Evaluate(el.ListElementValue(), newCtx)
				}

				if el.IsSymbol() {
					elBind := ctx.Lookup(el.SymbolElementValue())

					if elBind.IsElementBind() {
						el = elBind.GetElementValue()
					}
				}

				params = append(params, &el)
			}

			return bind.GetMethodValue().GetCall()(params, &ctx)
		}

		if bind.IsKeywordBind() {
			var params []*element.Element

			for _, param := range listElement.Children() {
				el := *param

				params = append(params, &el)
			}

			bindResult := bind.GetKeywordValue().GetCall()(params, &ctx)

			return bindResult
		}

		if bind.GetElementValue().IsFunction() {
			return EvaluateFun(*bind.GetElementValue().FunctionElementValue(), listElement, ctx)
		}
	}

	// anonymous function call

	if firstElement.IsFunction() {
		newCtx := ctx.Branch()
		return EvaluateFun(*firstElement.FunctionElementValue(), listElement, newCtx)
	}

	// not callable list

	var last element.Element

	last = *listElement

	for last.IsList() {
		result := element.ListElement{}

		for _, param := range listElement.Children() {
			el := *param

			if el.IsList() {
				newCtx := ctx.Branch()
				el = Evaluate(el.ListElementValue(), newCtx)
			}

			if el.IsSymbol() {
				elBind := ctx.Lookup(el.SymbolElementValue())

				if elBind.IsElementBind() {
					el = elBind.GetElementValue()
				}
			}

			result.Push(el)
		}

		last = *result.Last()
	}

	return last
}

func EvaluateFun(functionElement element.FunctionElement, listElement *element.ListElement, ctx language.Context) element.Element {
	var params []*element.Element

	for _, param := range listElement.Children() {
		el := *param

		if el.IsList() {
			newCtx := ctx.Branch()
			el = Evaluate(el.ListElementValue(), newCtx)
		}

		if el.IsSymbol() {
			elBind := ctx.Lookup(el.SymbolElementValue())

			if elBind.IsElementBind() {
				el = elBind.GetElementValue()
			}
		}

		params = append(params, &el)
	}

	functionCtx := ctx.Branch()

	children := functionElement.GetArguments().Children()

	for i, _ := range children {

		el := *children[i]
		functionCtx.Declare(*el.SymbolElementValue(), memory.NewElementBind(params[i+1]))
	}

	return Evaluate(functionElement.GetBody().ListElementValue(), functionCtx)
}
