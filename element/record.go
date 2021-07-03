package element

type RecordElement struct {
	label StringElement
	value Element
}

func NewRecordElement(label string, value Element) RecordElement {
	return RecordElement{
		label: NewStringElement(label),
		value: value,
	}
}

func (r RecordElement) Label() StringElement {
	return r.label
}

func (r RecordElement) Value() Element {
	return r.value
}

func (e RecordElement) IsList() bool {
	return false
}

func (s RecordElement) IsString() bool {
	return false
}

func (e RecordElement) IsNumber() bool {
	return false
}

func (e RecordElement) IsSymbol() bool {
	return false
}

func (e RecordElement) IsNull() bool {
	return false
}

func (e RecordElement) NumberValue() float64 {
	return 0
}

func (e RecordElement) StringValue() string {
	return e.label.Value + ": " + e.value.StringValue()
}

func (e RecordElement) Children() []*Element {
	return nil
}

func (e RecordElement) ListElementValue() *ListElement {
	return nil
}

func (e RecordElement) SymbolValue() string {
	return e.label.Value
}

func (e RecordElement) IsError() bool {
	return false
}

func (e RecordElement) BoolValue() bool {
	return true
}

func (e RecordElement) IsFunction() bool {
	return false
}

func (e RecordElement) FunctionElementValue() *FunctionElement {
	return &FunctionElement{hasName: false, body: ListElement{}, arguments: ListElement{}}
}

func (e RecordElement) SymbolElementValue() *SymbolElement {
	return &SymbolElement{Value: "bool"}
}

// object

func (o RecordElement) IsObject() bool {
	return false
}

func (o RecordElement) ObjectElementValue() *ObjectElement {
	return &ObjectElement{}
}

func (e RecordElement) StringElementValue() *StringElement {
	return &StringElement{Value: e.label.Value}
}

func (e RecordElement) NumberElementValue() *NumberElement {
	return &NumberElement{Value: 0}
}

func (e RecordElement) IsRecord() bool {
	return true
}

func (e RecordElement) RecordElementValue() *RecordElement {
	return &e
}

func (e RecordElement) IsPath() bool {
	return false
}

func (e RecordElement) PathElementValue() *PathElement {
	return &PathElement{}
}

func (e RecordElement) IsType() bool {
	return false
}

func (e RecordElement) TypeElementValue() *TypeElement {
	return &TypeElement{
		Type: TYPE_UNDEFINED,
	}
}

func (e RecordElement) IsBool() bool {
	return false
}
