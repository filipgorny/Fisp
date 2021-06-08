package main

type ExprList struct {
	elements []Element
	parent   *ExprList
}

func parse(tokens []LToken) ExprList {
	list := ExprList{}
	var currentList *ExprList
	currentList = &list

	for _, token := range tokens {
		if token.kind == LT_OPEN_LIST {
			newList := ExprList{parent: currentList}

			currentList = &newList
		} else if token.kind == LT_CLOSE_LIST {
			listElement := ListElement{}

			for _, el := range currentList.elements {
				listElement.elements = append(listElement.elements, el)
			}

			currentList = currentList.parent
			currentList.elements = append(currentList.elements, listElement)
		} else {
			if token.kind == LT_NUMBER {
				var el NumberElement
				el.numberValue = token.numberValue
				currentList.elements = append(currentList.elements, el)

			} else if token.kind == LT_SYMBOL {
				var el SymbolElement
				el.stringValue = token.stringValue

				currentList.elements = append(currentList.elements, el)
			}
		}
	}

	return list
}
