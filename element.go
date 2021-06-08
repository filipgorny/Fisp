package main

import "strconv"

type Element interface {
	IsList() bool
	NumberValue() float64
	StringValue() string
	ListElementValue() *ListElement
	Children() []*Element
}

// Symbol

type SymbolElement struct {
	stringValue string
}

func (s SymbolElement) IsList() bool {
	return false
}

func (s SymbolElement) NumberValue() float64 {
	return 0
}

func (s SymbolElement) StringValue() string {
	return s.stringValue
}

func (s SymbolElement) Children() []*Element {
	return nil
}

func (e SymbolElement) ListElementValue() *ListElement {
	return nil
}

// LIST

type ListElement struct {
	elements []Element
}

func (l ListElement) IsList() bool {
	return true
}

func (l ListElement) NumberValue() float64 {
	return 0
}

func (l ListElement) StringValue() string {
	var s = ""

	for _, element := range l.Children() {
		el := *element
		s = s + ", " + el.StringValue()
	}

	return s
}

func (l ListElement) Children() []*Element {
	var result []*Element

	for _, el := range l.elements {
		result = append(result, &el)
	}

	return result
}

func (l ListElement) ListElementValue() *ListElement {
	return &l
}

// Number

type NumberElement struct {
	numberValue float64
}

func (n NumberElement) IsList() bool {
	return false
}

func (n NumberElement) NumberValue() float64 {
	return n.numberValue
}

func (n NumberElement) StringValue() string {
	return strconv.FormatFloat(n.numberValue, 'E', -1, 64)
}

func (n NumberElement) Children() []*Element {
	return nil
}

func (n NumberElement) ListElementValue() *ListElement {
	return nil
}

// String

type StringElement struct {
	stringValue string
}

func (s StringElement) IsList() bool {
	return false
}

func (s StringElement) NumberValue() float64 {
	return 0
}

func (s StringElement) StringValue() string {
	return s.stringValue
}

func (s StringElement) Children() []*Element {
	return nil
}

func (s StringElement) ListElementValue() *ListElement {
	return nil
}
