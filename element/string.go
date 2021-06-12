package element

type StringElement struct {
	value string
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
	return s.value
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

func (s StringElement) IsError() bool {
	return false
}

func (e StringElement) BoolValue() bool {
	return len(e.value) > 0
}
