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

func (e ErrorElement) SymbolValue() string {
	return "ERROR"
}

func (e ErrorElement) IsError() bool {
	return true
}
