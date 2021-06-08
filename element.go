package main

import "strconv"

type Element interface {
	IsList() bool
	NumberValue() float64
	StringValue() string
	Children() *[]Element
}

// Symbol

type Symbol struct {
	stringValue string
}

func (s Symbol) IsList() bool {
	return false
}

func (s Symbol) NumberValue() float64 {
	return 0
}

func (s Symbol) StringValue() string {
	return s.stringValue
}

func (s Symbol) Children() *[]Element {
	return nil
}

// LIST

type List struct {
	children []Element
}

func (l List) IsList() bool {
	return true
}

func (l List) NumberValue() float64 {
	return 0
}

func (l List) StringValue() string {
	var s = ""

	for _, element := range l.children {
		s = s + ", " + element.StringValue()
	}

	return s
}

func (l List) Children() *[]Element {
	return &(l.children)
}

// Number

type Number struct {
	numberValue float64
}

func (n Number) IsList() bool {
	return false
}

func (n Number) NumberValue() float64 {
	return n.numberValue
}

func (n Number) StringValue() string {
	return strconv.FormatFloat(n.numberValue, 'E', -1, 64)
}

func (n Number) Children() *[]Element {
	return nil
}

// String

type String struct {
	stringValue string
}

func (s String) IsList() bool {
	return false
}

func (s String) NumberValue() float64 {
	return 0
}

func (s String) StringValue() string {
	return s.stringValue
}

func (s String) Children() *[]Element {
	return nil
}
