package element

type BoolElement struct {
	value bool
}

func NewBoolElement(value bool) BoolElement {
	return BoolElement{
		value: value,
	}
}

func (e BoolElement) IsList() bool {
	return false
}

func (e BoolElement) IsSymbol() bool {
	return false
}

func (e BoolElement) IsNull() bool {
	return false
}

func (e BoolElement) NumberValue() float64 {
	if e.value == true {
		return 1
	} else {
		return 0
	}
}

func (e BoolElement) StringValue() string {
	if e.value {
		return "true"
	} else {
		return "false"
	}
}

func (e BoolElement) Children() []*Element {
	return nil
}

func (e BoolElement) ListElementValue() *ListElement {
	return nil
}

func (e BoolElement) SymbolValue() string {
	if e.value {
		return "true"
	} else {
		return "false"
	}
}

func (e BoolElement) IsError() bool {
	return false
}

func (e BoolElement) BoolValue() bool {
	return e.value
}

func (e BoolElement) IsFunction() bool {
	return false
}

func (e BoolElement) FunctionElementValue() *FunctionElement {
	return &FunctionElement{hasName: false, body: ListElement{}, arguments: ListElement{}}
}

func (e BoolElement) SymbolElementValue() *SymbolElement {
	return &SymbolElement{Value: "bool"}
}
