package element

import "strconv"

type NumberElement struct {
	value float64
}

func (n NumberElement) IsList() bool {
	return false
}

func (n NumberElement) IsSymbol() bool {
	return true
}

func (n NumberElement) NumberValue() float64 {
	return n.value
}

func (n NumberElement) StringValue() string {
	return strconv.FormatFloat(n.value, 'f', 2, 64)
}

func (n NumberElement) Children() []*Element {
	return nil
}

func (n NumberElement) ListElementValue() *ListElement {
	return nil
}

func (n NumberElement) SymbolValue() string {
	return n.StringValue()
}

func (n NumberElement) IsError() bool {
	return false
}

func (e NumberElement) BoolValue() bool {
	return e.value > 0
}
