package element

import (
	"fmt"

	uuid "github.com/nu7hatch/gouuid"
)

type ObjectElement struct {
	OID    string
	Name   StringElement
	Fields []RecordElement
}

func NewObjectElement(name string, fields []RecordElement) ObjectElement {
	u, _ := uuid.NewV4()

	return ObjectElement{
		OID:    u.String(),
		Name:   NewStringElement(name),
		Fields: fields,
	}
}

func (e ObjectElement) IsList() bool {
	return false
}

func (s ObjectElement) IsString() bool {
	return false
}

func (e ObjectElement) IsNumber() bool {
	return false
}

func (e ObjectElement) IsSymbol() bool {
	return false
}

func (e ObjectElement) IsNull() bool {
	return false
}

func (e ObjectElement) NumberValue() float64 {
	return 1
}

func (e ObjectElement) StringValue() string {
	return fmt.Sprint("object<", e.Name, ">")
}

func (e ObjectElement) Children() []*Element {
	return nil
}

func (e ObjectElement) ListElementValue() *ListElement {
	return nil
}

func (e ObjectElement) SymbolValue() string {
	return ""
}

func (e ObjectElement) IsError() bool {
	return false
}

func (e ObjectElement) BoolValue() bool {
	return true
}

func (e ObjectElement) IsFunction() bool {
	return false
}

func (e ObjectElement) FunctionElementValue() *FunctionElement {
	return &FunctionElement{hasName: false, body: ListElement{}, arguments: ListElement{}}
}

func (e ObjectElement) SymbolElementValue() *SymbolElement {
	return &SymbolElement{Value: e.StringValue()}
}

// object

func (o ObjectElement) IsObject() bool {
	return true
}

func (o ObjectElement) ObjectElementValue() *ObjectElement {
	return &o
}

func (e ObjectElement) StringElementValue() *StringElement {
	return &StringElement{Value: "object"}
}

func (e ObjectElement) NumberElementValue() *NumberElement {
	return &NumberElement{Value: 1}
}

func (e ObjectElement) IsRecord() bool {
	return false
}

func (e ObjectElement) RecordElementValue() *RecordElement {
	return &RecordElement{value: e}
}

func (e ObjectElement) IsPath() bool {
	return false
}

func (e ObjectElement) PathElementValue() *PathElement {
	return &PathElement{}
}

func (e ObjectElement) IsType() bool {
	return false
}

func (e ObjectElement) TypeElementValue() *TypeElement {
	return &TypeElement{
		Type: TYPE_UNDEFINED,
	}
}
