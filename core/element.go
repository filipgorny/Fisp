package core

import (
	"log"
	"strconv"
)

type Element interface {
	IsList() bool
	NumberValue() float64
	StringValue() string
	ListElementValue() *ListElement
	Children() []*Element
	SymbolValue() string
}

// Symbol

type SymbolElement struct {
	Value string
}

func (s SymbolElement) IsList() bool {
	return false
}

func (s SymbolElement) NumberValue() float64 {
	return 0
}

func (s SymbolElement) StringValue() string {
	return s.Value
}

func (s SymbolElement) SymbolValue() string {
	return s.Value
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

func (l ListElement) SymbolValue() string {
	return l.StringValue()
}

func (l ListElement) Children() []*Element {
	var result []*Element

	for i, _ := range l.elements {
		result = append(result, &l.elements[i])
	}

	log.Print(result)

	return result
}

func (l ListElement) ListElementValue() *ListElement {
	return &l
}

func (l ListElement) All() []*Element {
	return l.Children()
}

// Number

type NumberElement struct {
	Value float64
}

func (n NumberElement) IsList() bool {
	return false
}

func (n NumberElement) NumberValue() float64 {
	return n.Value
}

func (n NumberElement) StringValue() string {
	return strconv.FormatFloat(n.Value, 'E', -1, 64)
}

func (n NumberElement) Children() []*Element {
	return nil
}

func (n NumberElement) ListElementValue() *ListElement {
	return nil
}

func (n NumberElement) SymbolValue() string {
	return n.StringValue()
}

// String

type StringElement struct {
	Value string
}

func (s StringElement) IsList() bool {
	return false
}

func (s StringElement) NumberValue() float64 {
	return 0
}

func (s StringElement) StringValue() string {
	return s.Value
}

func (s StringElement) Children() []*Element {
	return nil
}

func (s StringElement) ListElementValue() *ListElement {
	return nil
}

func (s StringElement) SymbolValue() string {
	return "\"" + s.StringValue() + "\""
}