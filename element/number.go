package element

import "strconv"

type NumberElement struct {
	Value float64
}

func (n NumberElement) IsList() bool {
	return false
}

func (s NumberElement) IsString() bool {
	return false
}

func (e NumberElement) IsNumber() bool {
	return true
}

func (n NumberElement) IsSymbol() bool {
	return true
}

func (n NumberElement) NumberValue() float64 {
	return n.Value
}

func (n NumberElement) StringValue() string {
	return strconv.FormatFloat(n.Value, 'f', 2, 64)
}

func (n NumberElement) Children() []*Element {
	return nil
}

func (n NumberElement) ListElementValue() *ListElement {
	return nil
}

func (n NumberElement) SymbolElementValue() *SymbolElement {
	return &SymbolElement{n.StringValue()}
}

func (n NumberElement) IsFunction() bool {
	return false
}

func (n NumberElement) FunctionElementValue() *FunctionElement {
	return &FunctionElement{hasName: false, body: ListElement{}, arguments: ListElement{}}
}

func (n NumberElement) IsError() bool {
	return false
}

func (e NumberElement) IsNull() bool {
	return false
}

func (e NumberElement) BoolValue() bool {
	return e.Value > 0
}

// object

func (o NumberElement) IsObject() bool {
	return false
}

func (o NumberElement) ObjectElementValue() *ObjectElement {
	return &ObjectElement{}
}

func (e NumberElement) StringElementValue() *StringElement {
	return &StringElement{Value: strconv.FormatFloat(e.Value, 'f', 2, 64)}
}

func (e NumberElement) NumberElementValue() *NumberElement {
	return &e
}

func (e NumberElement) IsRecord() bool {
	return false
}

func (e NumberElement) RecordElementValue() *RecordElement {
	return &RecordElement{value: e}
}

func (e NumberElement) IsPath() bool {
	return false
}

func (e NumberElement) PathElementValue() *PathElement {
	return &PathElement{}
}

func (e NumberElement) IsType() bool {
	return false
}

func (e NumberElement) TypeElementValue() *TypeElement {
	return &TypeElement{
		Type: TYPE_UNDEFINED,
	}
}

func (e NumberElement) IsBool() bool {
	return false
}
