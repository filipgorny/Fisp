package core

type ExprList struct {
	elements []Element
	parent   *ExprList
}

type Code interface {
	Elements() []Element
}

func (e ExprList) Elements() []Element {
	return e.elements
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
			listElement := ListElement{}

			for _, el := range currentList.elements {
				listElement.elements = append(listElement.elements, el)
			}

			currentList = currentList.parent
			currentList.elements = append(currentList.elements, listElement)
		} else {
			if token.kind == LT_NUMBER {
				var el NumberElement
				el.Value = token.numberValue
				currentList.elements = append(currentList.elements, el)

			} else if token.kind == LT_SYMBOL {
				var el SymbolElement
				el.Value = token.stringValue

				currentList.elements = append(currentList.elements, el)
			} else if token.kind == LT_STRING {
				var el StringElement
				el.Value = token.stringValue

				currentList.elements = append(currentList.elements, el)
			}
		}
	}

	return list
}
