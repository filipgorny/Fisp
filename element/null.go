package element

type NullElement struct {
}

func (n NullElement) IsList() bool {
	return false
}

func (s NullElement) IsString() bool {
	return false
}

func (e NullElement) IsNumber() bool {
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

// object

func (o NullElement) IsObject() bool {
	return false
}

func (o NullElement) ObjectElementValue() *ObjectElement {
	return &ObjectElement{}
}

func (e NullElement) StringElementValue() *StringElement {
	return &StringElement{Value: ""}
}

func (e NullElement) NumberElementValue() *NumberElement {
	return &NumberElement{Value: 0}
}

func (e NullElement) IsRecord() bool {
	return false
}

func (e NullElement) RecordElementValue() *RecordElement {
	return &RecordElement{value: e}
}

func (e NullElement) IsPath() bool {
	return false
}

func (e NullElement) PathElementValue() *PathElement {
	return &PathElement{}
}
