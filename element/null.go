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

func (n NullElement) SymbolElementValue() *SymbolElement {
	return &SymbolElement{Value: "null"}
}

func (n NullElement) IsFunction() bool {
	return false
}

func (n NullElement) FunctionElementValue() *FunctionElement {
	return &FunctionElement{hasName: false, body: ListElement{}, arguments: ListElement{}}
}

func (n NullElement) IsError() bool {
	return false
}

func (e NullElement) IsNull() bool {
	return true
}

func (e NullElement) BoolValue() bool {
	return false
}
