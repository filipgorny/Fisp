package runtime

import (
	"wxl/element"
	"wxl/logging"
)

type Context struct {
	parent *Context
	memory map[element.SymbolElement]Bind
	log    *logging.Log
}

func NewContext(env Environment, log *logging.Log) Context {
	ctx := Context{
		memory: make(map[element.SymbolElement]Bind),
		log:    log,
	}

	for _, m := range env.Methods {
		ctx.Declare(m.Symbol, MethodBind{method: *m})
	}

	return ctx
}

func (ctx *Context) Branch() Context {
	return Context{parent: ctx, log: ctx.log, memory: make(map[element.SymbolElement]Bind)}
}

func (ctx *Context) Declare(s element.SymbolElement, bind Bind) {
	if ctx.parent != nil {
		ctx.parent.Declare(s, bind)
	} else {
		ctx.memory[s] = bind
	}
}

func (ctx *Context) Lookup(s element.Element) Bind {
	if val, ok := ctx.memory[*s.SymbolElementValue()]; ok {
		ctx.Log("Found lookup value for `" + s.StringValue() + "` which is `" + val.GetElementValue().StringValue() + "`.")
		return val
	}

	if nil != ctx.parent {
		return ctx.parent.Lookup(s)
	}

	return NullBind{}
}

func (ctx Context) Log(content interface{}) {
	ctx.log.Info(content)
}

func (ctx Context) Error(content interface{}) {
	ctx.log.Error(content)
}
