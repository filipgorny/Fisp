package core

import (
	"strings"
	"wxl/element"
)

type ExprList struct {
	Elements []element.Element
	parent   *ExprList
}

type CurrentRecord struct {
	opened bool
	label  string
	isList bool
}

func Parse(tokens []LToken) ExprList {
	list := ExprList{}
	var currentList *ExprList
	currentList = &list

	currentRecord := CurrentRecord{}

	// go func() {
	for _, token := range tokens {
		if token.kind == LT_OPEN_LIST {
			newList := ExprList{parent: currentList}

			currentList = &newList

			if currentRecord.opened {
				currentRecord.isList = true
			}
		} else {
			var el element.Element

			if token.kind == LT_CLOSE_LIST {
				listElement := element.ListElement{}

				for _, el := range currentList.Elements {
					listElement.Elements = append(listElement.Elements, el)
				}

				currentList = currentList.parent

				el = listElement

				if !currentRecord.isList {
					currentRecord.opened = false
				}
			} else if token.kind == LT_NUMBER {
				el = element.NumberElement{Value: token.numberValue}
			} else if token.kind == LT_SYMBOL {
				if token.stringValue == "true" {
					el = element.NewBoolElement(true)
				} else if token.stringValue == "false" {
					el = element.NewBoolElement(false)
				} else {
					el = element.SymbolElement{Value: token.stringValue}
				}
			} else if token.kind == LT_STRING {
				el = element.StringElement{Value: token.stringValue}
			} else if token.kind == LT_PATH {
				el = element.NewPathElement(strings.Split(token.stringValue, "."))
			} else if token.kind == LT_TYPE {
				el = element.NewTypeElement(strings.Replace(token.stringValue, "^", "", 1))
			} else if token.kind == LT_LABEL {
				currentRecord.label = token.stringValue
				currentRecord.opened = true
			}

			if el != nil {
				if currentRecord.opened {
					el = element.NewRecordElement(currentRecord.label, el)

					currentRecord.opened = false
				}

				currentList.Elements = append(currentList.Elements, el)
			}
		}
	}
	// }()

	return list
}
