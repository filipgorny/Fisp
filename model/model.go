package model

import (
	"wxl/object"
)

type Model struct {
	name   string
	fields []*Field
}

func (o Model) Class() object.Class {
	return object.Class{Name: "model"}
}

func (o Model) Fields() []*Field {
	return o.fields
}

func (o Model) Name() string {
	return o.name
}

func (o Model) Field(name string) *Field {
	for _, f := range o.fields {
		if f.name == name {
			return f
		}
	}

	return &UndeclaredField
}

func NewModel(name string, fields []*Field) Model {
	return Model{
		name:   name,
		fields: fields,
	}
}
