package element

type NullElement struct {
}

func (n NullElement) IsList() bool {
	return false
}

func (n NullElement) IsSymbol() bool {
	return false
}

func (n NullElement) NumberValue() float64 {
	return 0
}

func (n NullElement) StringValue() string {
	return "null"
}

func (n NullElement) Children() []*Element {
	return nil
}

func (n NullElement) ListElementValue() *ListElement {
	return nil
}

func (n NullElement) SymbolValue() string {
	return "null"
}
