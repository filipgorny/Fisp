package main

import (
	"log"
)

type List interface {
	Pop() *Element
	Push(element Element)
	Last() *Element
}

func (list *ListElement) Pop() *Element {
	if len(list.elements) == 0 {
		log.Fatal("List is empty, cannot pop() from it.")
	}

	last := list.elements[len(list.elements)-1]

	list.elements = list.elements[:len(list.elements)-1]

	return &last
}

func (list *ListElement) Push(el Element) {
	list.elements = append(list.elements, el)
}

func (list ListElement) Last() *Element {
	return &list.elements[len(list.elements)-1]
}
