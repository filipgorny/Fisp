package model

import (
	"wxl/element"
)

type Type struct {
	Name     string
	Validate func(d element.Element) bool
}

func NewType(name string, validate func(d element.Element) bool) Type {
	return Type{
		Name:     name,
		Validate: validate,
	}
}

var TypeString = NewType("string", func(d element.Element) bool { return d.IsString() })
var TypeNumber = NewType("number", func(d element.Element) bool { return d.IsNumber() })
var TypeBool = NewType("bool", func(d element.Element) bool { return d.StringValue() == "false" || d.StringValue() == "true" })
var TypeFunction = NewType("function", func(d element.Element) bool { return d.IsFunction() })
var TypeObject = NewType("object", func(d element.Element) bool { return d.IsObject() })
var TypeVoid = NewType("void", func(d element.Element) bool { return d.IsNull() })

var Types = []Type{TypeString, TypeNumber, TypeBool, TypeFunction, TypeObject, TypeVoid}

func ValidTypeName(s string) bool {
	for _, t := range Types {
		if s == t.Name {
			return true
		}
	}

	return false
}

func TypeByName(s string) Type {
	for _, t := range Types {
		if s == t.Name {
			return t
		}
	}

	return TypeVoid
}
