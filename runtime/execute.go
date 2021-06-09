package runtime

import (
	"fmt"
	"wxl/core"
	"wxl/language"
)

func Run(spec language.Spec, codeTree core.ExprList) {
	for _, element := range codeTree.Elements() {

		if element.IsList() {
			firstElement := *element.ListElementValue().First()

			for _, def := range spec.Definitions {
				if def.Symbol.Value == firstElement.StringValue() {
					params := element.Children()
					call := def.Call(params)

					fmt.Print("PIERWSZY WYNIK")
					fmt.Print(call)
				}
			}
		}
	}
}
