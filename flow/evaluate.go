package flow

import "wxl/element"

func evaluate(listElement *element.ListElement, ctx *Context) element.Element {
	firstElement := listElement.First()

	for _, def := range spec.Definitions {
		if def.Symbol.Value == (*firstElement).StringValue() {
			var params []*element.Element

			for _, param := range listElement.Children() {
				el := *param

				if el.IsList() {
					newCtx := ctx.Branch()
					el = evaluate(el.ListElementValue(), &newCtx)
				}

				params = append(params, &el)
			}

			return def.Call(params, ctx)
		}
	}

	return element.NullElement{}
}
