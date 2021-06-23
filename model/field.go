package model

import (
	"wxl/object"
)

type Field struct {
	Type    Type
	name    string
	Defined bool
}

func (f Field) Class() object.Class {
	return object.Class{Name: "field"}
}

func NewField(name string, fieldType Type) Field {
	return Field{name: name, Type: fieldType, Defined: true}
}

var UndeclaredField = Field{Type: TypeVoid, name: "undeclared", Defined: false}
