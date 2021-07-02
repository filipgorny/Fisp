package stack

import (
	"wxl/element"
	"wxl/language"
	"wxl/logging"
	"wxl/memory"
	"wxl/object"
	"wxl/resource"
	"wxl/runtime"
)

type SystemNodeResource struct {
}

func (r SystemNodeResource) Class() object.Class {
	return object.Class{Name: "system-resource"}
}

func (r SystemNodeResource) Get(s resource.Selector) resource.Resource {
	return r
}

func (r SystemNodeResource) Put(selector resource.Selector, resource resource.Resource) {

}

type SystemContext struct {
	parent *SystemContext
	memory map[element.SymbolElement]language.Bind
	log    *logging.Log
	root   SystemNodeResource
}

func NewContext(env runtime.Environment, log *logging.Log) SystemContext {
	ctx := SystemContext{
		memory: make(map[element.SymbolElement]language.Bind),
		log:    log,
	}

	for _, m := range env.Methods {
		ctx.Declare(m.Symbol, memory.NewMethodBind(m))
	}

	for _, k := range env.Keywords {
		ctx.Declare(k.Symbol, memory.NewKeywordBind(k))
	}

	return ctx
}

func (ctx SystemContext) Branch() language.Context {
	return SystemContext{parent: &ctx, log: ctx.log, memory: make(map[element.SymbolElement]language.Bind)}
}

func (ctx SystemContext) Declare(s element.SymbolElement, bind language.Bind) {
	if ctx.parent != nil {
		ctx.parent.Declare(s, bind)
	} else {
		ctx.memory[s] = bind
	}
}

func (ctx SystemContext) Lookup(s element.Element) language.Bind {
	if val, ok := ctx.memory[*s.SymbolElementValue()]; ok {
		return val
	}

	if nil != ctx.parent {
		return ctx.parent.Lookup(s)
	}

	return memory.NullBind{}
}

func (ctx SystemContext) Log(content interface{}) {
	ctx.log.Info(content)
}

func (ctx SystemContext) Error(content interface{}) {
	ctx.log.Error(content)
}

func (ctx SystemContext) Get(s resource.Selector) resource.Resource {
	return ctx.root
}

func (ctx SystemContext) Class() object.Class {
	return object.Class{Name: "system-context"}
}
