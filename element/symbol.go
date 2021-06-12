package element

type SymbolElement struct {
	Value string
}

func (s SymbolElement) IsList() bool {
	return false
}

func (s SymbolElement) IsSymbol() bool {
	return true
}

func (s SymbolElement) NumberValue() float64 {
	return 0
}

func (s SymbolElement) StringValue() string {
	return s.Value
}

func (s SymbolElement) SymbolValue() string {
	return s.Value
}

func (s SymbolElement) Children() []*Element {
	return nil
}

func (e SymbolElement) ListElementValue() *ListElement {
	return nil
}

func (e SymbolElement) IsError() bool {
	return false
}

func (e SymbolElement) BoolValue() bool {
	return true
}
