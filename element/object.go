package element

import (
	uuid "github.com/nu7hatch/gouuid"
)

type ObjectElement struct {
	OID          string
	Name         StringElement
	Properties   []RecordElement
	nullProperty Element
}

func NewObjectElement(name string, properties []RecordElement) ObjectElement {
	u, _ := uuid.NewV4()

	return ObjectElement{
		OID:          u.String(),
		Name:         NewStringElement(name),
		Properties:   properties,
		nullProperty: NullElement{},
	}
}

func (e ObjectElement) IsA(name string) bool {
	return e.Name.Value == name
}

func (e *ObjectElement) Get(name string) Element {
	for _, property := range e.Properties {
		if property.label.Value == name {
			return property.value
		}
	}

	return e.nullProperty
}

func (e *ObjectElement) Set(name string, value Element) {
	newProperties := []RecordElement{}

	found := false

	for _, property := range e.Properties {
		if property.label == e.Name {
			newProperties = append(newProperties, NewRecordElement(name, value))

			found = true
		} else {
			newProperties = append(newProperties, property)
		}
	}

	if !found {
		newProperties = append(newProperties, NewRecordElement(name, value))
	}

	e.Properties = newProperties
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
	return e.Serialize().StringValue()
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
	str := NewStringElement(e.StringValue())

	return &str
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

func (e ObjectElement) IsBool() bool {
	return false
}

func (e *ObjectElement) Serialize() ListElement {
	resultList := NewListElement()

	resultList.Push(NewSymbolElement("object"))
	resultList.Push(e.Name)

	propertiesList := NewListElement()

	for _, property := range e.Properties {
		propertiesList.Push(NewRecordElement(property.label.Value, property.value.StringElementValue()))
	}

	resultList.Push(propertiesList)

	return resultList
}
