package element

type BoolElement struct {
	value bool
}

func (e BoolElement) IsList() bool {
	return false
}

func (e BoolElement) IsSymbol() bool {
	return false
}

func (e BoolElement) NumberValue() float64 {
	return 1
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
