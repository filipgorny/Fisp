package core

import "rpl/element"

type List interface {
	Pop() *element.Element
	Push(element element.Element)
	Last() *element.Element
	All() []*element.Element
	First() *element.Element
}
