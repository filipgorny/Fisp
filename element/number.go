package element

import "strconv"

type NumberElement struct {
	Value float64
}

func (n NumberElement) IsList() bool {
	return false
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
