package element

type ErrorElement struct {
	Message string
}

func (e ErrorElement) IsList() bool {
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
