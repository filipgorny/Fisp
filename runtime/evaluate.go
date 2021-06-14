package runtime

import "wxl/element"

func Evaluate(listElement *element.ListElement, ctx *Context) element.Element {
	firstElement := *listElement.First()

	if firstElement.IsSymbol() {
		bind := ctx.Lookup(firstElement.SymbolElementValue())

		if bind.IsMethodBind() {
			ctx.Log("method bind: " + bind.GetMethodValue().Symbol.StringValue())
			var params []*element.Element

			for _, param := range listElement.Children() {
				el := *param

				if el.IsList() {
					ctx.Log("Branching list: " + el.StringValue())
					newCtx := ctx.Branch()
					el = Evaluate(el.ListElementValue(), &newCtx)
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

			return bind.GetMethodValue().Call(params, ctx)
		}
	}

	return element.NullElement{}
}
