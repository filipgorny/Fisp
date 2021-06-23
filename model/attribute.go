package model

import "wxl/element"

type Attribute struct {
	Field Field
	Value element.Element
}

func NewAttribute(field Field, value element.Element) Attribute {
	return Attribute{Field: field, Value: value}
}

var UndefinedAttribute = NewAttribute(UndeclaredField, element.NewStringElement("undefined"))
