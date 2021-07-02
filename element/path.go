package element

import "wxl/resource"

type PathElement struct {
	path resource.PathSelector
}

func NewPathElement(path resource.PathSelector) PathElement {
	return PathElement{
		path: path,
	}
}

func (p PathElement) Path() resource.PathSelector {
	return p.path
}

func (e PathElement) IsList() bool {
	return false
}

func (s PathElement) IsString() bool {
	return false
}

func (e PathElement) IsNumber() bool {
	return false
}

func (e PathElement) IsSymbol() bool {
	return false
}

func (e PathElement) IsNull() bool {
	return false
}

func (e PathElement) NumberValue() float64 {
	return 1
}

func (e PathElement) StringValue() string {
	return e.path.String()
}

func (e PathElement) Children() []*Element {
	return nil
}

func (e PathElement) ListElementValue() *ListElement {
	return nil
}

func (e PathElement) SymbolValue() string {
	return e.path.String()
}

func (e PathElement) IsError() bool {
	return false
}

func (e PathElement) BoolValue() bool {
	return true
}

func (e PathElement) IsFunction() bool {
	return false
}

func (e PathElement) FunctionElementValue() *FunctionElement {
	return &FunctionElement{hasName: false, body: ListElement{}, arguments: ListElement{}}
}

func (e PathElement) SymbolElementValue() *SymbolElement {
	return &SymbolElement{Value: e.StringValue()}
}

// object

func (o PathElement) IsObject() bool {
	return false
}

func (o PathElement) ObjectElementValue() *ObjectElement {
	return &ObjectElement{}
}

func (e PathElement) StringElementValue() *StringElement {
	return &StringElement{Value: e.StringValue()}
}

func (e PathElement) NumberElementValue() *NumberElement {
	return &NumberElement{Value: 1}
}

func (e PathElement) IsRecord() bool {
	return false
}

func (e PathElement) RecordElementValue() *RecordElement {
	return &RecordElement{value: e}
}

func (e PathElement) IsPath() bool {
	return true
}

func (e PathElement) PathElementValue() *PathElement {
	return &e
}

func (e PathElement) IsType() bool {
	return false
}

func (e PathElement) TypeElementValue() *TypeElement {
	return &TypeElement{
		Type: TYPE_UNDEFINED,
	}
}
