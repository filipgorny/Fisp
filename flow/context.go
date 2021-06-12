package flow

import (
	"rpl/element"
	"rpl/language"
)

type Context struct {
	parent *Context
	memory map[string]element.Element
	spec   language.Spec
}

func NewContext(env Environment) Context {
	return Context{
		memory: make(map[string]element.Element),
	}
}

func (ctx Context) Branch() Context {
	return Context{parent: &ctx}
}

func (ctx Context) Declare(s element.SymbolElement, value element.Element) {
	ctx.memory[s.StringValue()] = value
}

func (ctx Context) Lookup(s element.SymbolElement) element.Element {
	if val, ok := ctx.memory[s.StringValue()]; ok {
		return val
	}

	return element.ErrorElement{Message: "undeclared variable: " + s.StringValue()}
}
