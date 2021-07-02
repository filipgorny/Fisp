package element

import (
	"strconv"
	"wxl/resource"
)

type StringElement struct {
	Value string
}

func NewStringElement(value string) StringElement {
	return StringElement{
		Value: value,
	}
}

func (s StringElement) IsList() bool {
	return false
}

func (s StringElement) IsString() bool {
	return true
}

func (e StringElement) IsNumber() bool {
	return false
}

func (s StringElement) IsSymbol() bool {
	return true
}

func (s StringElement) NumberValue() float64 {
	num, _ := strconv.ParseFloat(s.Value, 64)

	return num
}

func (s StringElement) StringValue() string {
	return s.Value
}

func (s StringElement) Children() []*Element {
	return nil
}

func (s StringElement) ListElementValue() *ListElement {
	return nil
}

func (s StringElement) SymbolElementValue() *SymbolElement {
	return &SymbolElement{s.StringValue()}
}

func (s StringElement) IsFunction() bool {
	return false
}

func (s StringElement) FunctionElementValue() *FunctionElement {
	return &FunctionElement{hasName: false, body: ListElement{}, arguments: ListElement{}}
}

func (s StringElement) IsError() bool {
	return false
}

func (e StringElement) IsNull() bool {
	return false
}

func (e StringElement) BoolValue() bool {
	return len(e.Value) > 0
}

// object

func (o StringElement) IsObject() bool {
	return false
}

func (o StringElement) ObjectElementValue() *ObjectElement {
	return &ObjectElement{}
}

func (e StringElement) StringElementValue() *StringElement {
	return &e
}

func (e StringElement) NumberElementValue() *NumberElement {
	num, _ := strconv.ParseFloat(e.Value, 64)

	return &NumberElement{Value: num}
}

func (e StringElement) IsRecord() bool {
	return false
}

func (e StringElement) RecordElementValue() *RecordElement {
	return &RecordElement{value: e}
}

func (e StringElement) IsPath() bool {
	return false
}

func (e StringElement) PathElementValue() *PathElement {
	return &PathElement{path: resource.NewPath([]string{e.StringValue()})}
}

func (e StringElement) IsType() bool {
	return false
}

func (e StringElement) TypeElementValue() *TypeElement {
	return &TypeElement{
		Type: TYPE_UNDEFINED,
	}
}
