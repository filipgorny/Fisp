package element

import (
	"rpl/environment"
)

type SystemMethodElement struct {
	Method environment.Method
}

func (s SystemMethodElement) IsList() bool {
	return false
}

func (s SystemMethodElement) IsSystemMethod() bool {
	return true
}

func (s SystemMethodElement) NumberValue() float64 {
	return 0
}

func (s SystemMethodElement) StringValue() string {
	return s.Method.Symbol.StringValue()
}

func (s SystemMethodElement) SystemMethodValue() string {
	return s.Value
}

func (s SystemMethodElement) Children() []*Element {
	return nil
}

func (e SystemMethodElement) ListElementValue() *ListElement {
	return nil
}

func (e SystemMethodElement) IsError() bool {
	return false
}
