package core

import "wxl/element"

type ExprList struct {
	Elements []element.Element
	parent   *ExprList
}

func Parse(tokens []LToken) ExprList {
	list := ExprList{}
	var currentList *ExprList
	currentList = &list

	for _, token := range tokens {
		if token.kind == LT_OPEN_LIST {
			newList := ExprList{parent: currentList}

			currentList = &newList
		} else if token.kind == LT_CLOSE_LIST {
			listElement := element.ListElement{}

			for _, el := range currentList.Elements {
				listElement.Elements = append(listElement.Elements, el)
			}

			currentList = currentList.parent
			currentList.Elements = append(currentList.Elements, listElement)
		} else {
			if token.kind == LT_NUMBER {
				var el element.NumberElement
				el.Value = token.numberValue
				currentList.Elements = append(currentList.Elements, el)

			} else if token.kind == LT_SYMBOL {
				var el element.SymbolElement
				el.Value = token.stringValue

				currentList.Elements = append(currentList.Elements, el)
			} else if token.kind == LT_STRING {
				var el element.StringElement
				el.Value = token.stringValue

				currentList.Elements = append(currentList.Elements, el)
			}
		}
	}

	return list
}
