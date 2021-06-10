package element

type StringElement struct {
	Value string
}

func (s StringElement) IsList() bool {
	return false
}

func (s StringElement) IsSymbol() bool {
	return true
}

func (s StringElement) NumberValue() float64 {
	return 0
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

func (s StringElement) SymbolValue() string {
	return "\"" + s.StringValue() + "\""
}
