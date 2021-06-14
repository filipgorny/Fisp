package stack

import (
	"wxl/element"
	"wxl/language"
	"wxl/logging"
	"wxl/memory"
	"wxl/runtime"
)

type Context struct {
	parent *Context
	memory map[element.SymbolElement]language.Bind
	log    *logging.Log
}

func NewContext(env runtime.Environment, log *logging.Log) Context {
	ctx := Context{
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

func (ctx Context) Branch() language.Context {
	return Context{parent: &ctx, log: ctx.log, memory: make(map[element.SymbolElement]language.Bind)}
}

func (ctx Context) Declare(s element.SymbolElement, bind language.Bind) {
	if ctx.parent != nil {
		ctx.parent.Declare(s, bind)
	} else {
		ctx.memory[s] = bind
	}
}

func (ctx Context) Lookup(s element.Element) language.Bind {
	if val, ok := ctx.memory[*s.SymbolElementValue()]; ok {
		ctx.Log("Found lookup value for `" + s.StringValue() + "` which is `" + val.GetElementValue().StringValue() + "`.")
		return val
	}

	if nil != ctx.parent {
		return ctx.parent.Lookup(s)
	}

	return memory.NullBind{}
}

func (ctx Context) Log(content interface{}) {
	ctx.log.Info(content)
}

func (ctx Context) Error(content interface{}) {
	ctx.log.Error(content)
}
