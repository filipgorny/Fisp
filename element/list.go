package element

import "log"

type ListElement struct {
	Elements []Element
}

func (l ListElement) IsList() bool {
	return true
}

func (l ListElement) NumberValue() float64 {
	return 0
}

func (l ListElement) StringValue() string {
	var s = ""

	for _, element := range l.Children() {
		el := *element
		s = s + ", " + el.StringValue()
	}

	return s
}

func (e ListElement) IsNumber() bool {
	return false
}

func (l ListElement) SymbolElementValue() *SymbolElement {
	return &SymbolElement{Value: "list"}
}

func (l ListElement) IsFunction() bool {
	return false
}

func (l ListElement) FunctionElementValue() *FunctionElement {
	return &FunctionElement{hasName: false, body: ListElement{}, arguments: ListElement{}}
}

func (l ListElement) IsSymbol() bool {
	return true
}

func (s ListElement) IsString() bool {
	return false
}

func (e ListElement) IsNull() bool {
	return false
}

func (l ListElement) Children() []*Element {
	var result []*Element

	for i := range l.Elements {
		result = append(result, &l.Elements[i])
	}

	return result
}

func (l ListElement) ListElementValue() *ListElement {
	return &l
}

func (l ListElement) All() []*Element {
	return l.Children()
}

//

func (list *ListElement) Pop() *Element {
	if len(list.Elements) == 0 {
		log.Fatal("List is empty, cannot pop() from it.")
	}

	last := list.Elements[len(list.Elements)-1]

	list.Elements = list.Elements[:len(list.Elements)-1]

	return &last
}

func (list *ListElement) Push(el Element) {
	list.Elements = append(list.Elements, el)
}

func (list ListElement) Last() *Element {
	if len(list.Elements) < 1 {
		var null Element

		null = NullElement{}

		return &null
	}

	return &list.Elements[len(list.Elements)-1]
}

func (list ListElement) First() *Element {
	if len(list.Elements) < 1 {
		var null Element

		null = NullElement{}

		return &null
	}

	return &list.Elements[0]
}

func (list ListElement) IsError() bool {
	return false
}

func (e ListElement) BoolValue() bool {
	return len(e.Elements) > 0
}

// object

func (o ListElement) IsObject() bool {
	return false
}

func (o ListElement) ObjectElementValue() *ObjectElement {
	return &ObjectElement{}
}

func (e ListElement) StringElementValue() *StringElement {
	return &StringElement{Value: "list"}
}

func (e ListElement) NumberElementValue() *NumberElement {
	return &NumberElement{Value: 0}
}

func (e ListElement) IsRecord() bool {
	return false
}

func (e ListElement) RecordElementValue() *RecordElement {
	return &RecordElement{value: e}
}

func (e ListElement) IsPath() bool {
	return false
}

func (e ListElement) PathElementValue() *PathElement {
	return &PathElement{}
}
