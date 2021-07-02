package element

type SymbolElement struct {
	Value string
}

func (s SymbolElement) IsList() bool {
	return false
}

func (s SymbolElement) IsString() bool {
	return false
}

func (s SymbolElement) IsSymbol() bool {
	return true
}

func (e SymbolElement) IsNumber() bool {
	return false
}

func (s SymbolElement) NumberValue() float64 {
	return 0
}

func (s SymbolElement) StringValue() string {
	return s.Value
}

func (s SymbolElement) SymbolElementValue() *SymbolElement {
	return &s
}

func (s SymbolElement) IsFunction() bool {
	return false
}

func (s SymbolElement) FunctionElementValue() *FunctionElement {
	return &FunctionElement{hasName: false, body: ListElement{}, arguments: ListElement{}}
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

func (e SymbolElement) IsNull() bool {
	return false
}

func (e SymbolElement) BoolValue() bool {
	return true
}

// object

func (o SymbolElement) IsObject() bool {
	return false
}

func (o SymbolElement) ObjectElementValue() *ObjectElement {
	return &ObjectElement{}
}

func (e SymbolElement) StringElementValue() *StringElement {
	return &StringElement{Value: e.Value}
}

func (e SymbolElement) NumberElementValue() *NumberElement {
	return &NumberElement{Value: 1}
}

func (e SymbolElement) IsRecord() bool {
	return false
}

func (e SymbolElement) RecordElementValue() *RecordElement {
	return &RecordElement{value: e}
}

func (e SymbolElement) IsPath() bool {
	return false
}

func (e SymbolElement) PathElementValue() *PathElement {
	return &PathElement{}
}

func (e SymbolElement) IsType() bool {
	return false
}

func (e SymbolElement) TypeElementValue() *TypeElement {
	return &TypeElement{
		Type: TYPE_UNDEFINED,
	}
}
