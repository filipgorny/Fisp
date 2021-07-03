package web

import (
	"wxl/element"
	"wxl/language"
	"wxl/logging"
	"wxl/memory"
	"wxl/resource"
	"wxl/runtime"
)

type WebContext struct {
	parent *WebContext
	memory map[element.SymbolElement]language.Bind
	log    *logging.Log
	root   WebNodeResource
}

func NewContext(env runtime.Environment, log *logging.Log) WebContext {
	ctx := WebContext{
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

func (ctx WebContext) Declare(s element.SymbolElement, bind language.Bind) {
	if ctx.parent != nil {
		ctx.parent.Declare(s, bind)
	} else {
		ctx.memory[s] = bind
	}
}

func (ctx WebContext) Lookup(s element.Element) language.Bind {
	if val, ok := ctx.memory[*s.SymbolElementValue()]; ok {
		return val
	}

	if nil != ctx.parent {
		return ctx.parent.Lookup(s)
	}

	return memory.NullBind{}
}

func (ctx WebContext) Put(p element.PathElement, value resource.Resource) {
	ctx.root.Put(p, value)
}

func (ctx WebContext) Get(p element.PathElement) resource.Resource {
	return ctx.root.Get(p)
}

func (ctx WebContext) Branch() *language.Context {
	var newCtx language.Context = ctx

	return &newCtx
}
