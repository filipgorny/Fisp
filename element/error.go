package element

import "strings"

type ErrorElement struct {
	Message string
}

func NewErrorElement(messages ...string) ErrorElement {
	return ErrorElement{
		Message: strings.Join(messages, " "),
	}
}

func (e ErrorElement) IsList() bool {
	return false
}

func (s ErrorElement) IsString() bool {
	return false
}

func (e ErrorElement) IsNumber() bool {
	return false
}

func (e ErrorElement) IsSymbol() bool {
	return false
}

func (e ErrorElement) NumberValue() float64 {
	return -1
}

func (e ErrorElement) StringValue() string {
	return e.Message
}

func (e ErrorElement) Children() []*Element {
	return nil
}

func (e ErrorElement) ListElementValue() *ListElement {
	return nil
}

func (e ErrorElement) SymbolElementValue() *SymbolElement {
	return &SymbolElement{Value: "error"}
}

func (e ErrorElement) IsError() bool {
	return true
}

func (e ErrorElement) IsNull() bool {
	return false
}

func (e ErrorElement) BoolValue() bool {
	return false
}

func (e ErrorElement) IsFunction() bool {
	return false
}

func (e ErrorElement) FunctionElementValue() *FunctionElement {
	return &FunctionElement{hasName: false, body: ListElement{}, arguments: ListElement{}}
}

// object

func (o ErrorElement) IsObject() bool {
	return false
}

func (o ErrorElement) ObjectElementValue() *ObjectElement {
	return &ObjectElement{}
}

func (e ErrorElement) StringElementValue() *StringElement {
	return &StringElement{Value: e.Message}
}

func (e ErrorElement) NumberElementValue() *NumberElement {
	return &NumberElement{Value: -1}
}

func (e ErrorElement) IsRecord() bool {
	return false
}

func (e ErrorElement) RecordElementValue() *RecordElement {
	return &RecordElement{value: e}
}

func (e ErrorElement) IsPath() bool {
	return false
}

func (e ErrorElement) PathElementValue() *PathElement {
	return &PathElement{}
}

func (e ErrorElement) IsType() bool {
	return false
}

func (e ErrorElement) TypeElementValue() *TypeElement {
	return &TypeElement{
		Type: TYPE_UNDEFINED,
	}
}

func (e ErrorElement) IsBool() bool {
	return false
}
