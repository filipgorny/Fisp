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

func (s StringElement) SymbolElementValue() *SymbolElement {
	return &SymbolElement{s.StringValue()}
}

func (s StringElement) IsFunction() bool {
	return false
}

func (s StringElement) FunctionElementValue() *FunctionElement {
	return &FunctionElement{hasName: false, body: ListElement{}, arguments: ListElement{}}
}

func (s StringElement) IsError() bool {
	return false
}

func (e StringElement) IsNull() bool {
	return false
}

func (e StringElement) BoolValue() bool {
	return len(e.Value) > 0
}