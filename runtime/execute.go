package runtime

import (
	"log"
	"rpl/core"
	"rpl/element"
	"rpl/language"
)

func evaluateList(spec language.Spec, listElement *element.ListElement) element.Element {
	firstElement := listElement.First()

	for _, def := range spec.Definitions {
		if def.Symbol.Value == (*firstElement).StringValue() {
			var params []*element.Element

			for _, param := range listElement.Children() {
				el := *param

				if el.IsList() {
					el = evaluateList(spec, el.ListElementValue())
				}

				params = append(params, &el)
			}

			return def.Call(params).Element
		}
	}

	return element.NullElement{}
}

func Run(spec language.Spec, codeTree core.ExprList) {
	for _, element := range codeTree.Elements {
		if element.IsList() {
			result := evaluateList(spec, element.ListElementValue())

			log.Println("Result: " + result.StringValue())
		}
	}
}
