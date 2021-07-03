package stack

import (
	"wxl/element"
	"wxl/language"
	"wxl/logging"
	"wxl/memory"
	"wxl/resource"
	"wxl/runtime"
)

var SystemResourceMemory map[string]resource.Resource

type SystemNodeResource struct {
}

func (r SystemNodeResource) Get(path element.PathElement) resource.Resource {
	if SystemResourceMemory[path.String()] == nil {
		return element.NullElement{}
	}

	return SystemResourceMemory[path.String()]
}

func (r SystemNodeResource) Put(path element.PathElement, resource resource.Resource) {
	SystemResourceMemory[path.String()] = resource
}

type SystemContext struct {
	parent *SystemContext
	memory map[element.SymbolElement]language.Bind
	log    *logging.Log
	root   SystemNodeResource
}

func NewContext(env runtime.Environment, log *logging.Log) SystemContext {
	node := SystemNodeResource{}

	ctx := SystemContext{
		memory: make(map[element.SymbolElement]language.Bind),
		log:    log,
		root:   node,
	}

	SystemResourceMemory = make(map[string]resource.Resource)

	for _, m := range env.Methods {
		ctx.Declare(m.Symbol, memory.NewMethodBind(m))
	}

	for _, k := range env.Keywords {
		ctx.Declare(k.Symbol, memory.NewKeywordBind(k))
	}

	return ctx
}

func (ctx SystemContext) Branch() *language.Context {
	var newCtx language.Context = SystemContext{parent: &ctx, log: ctx.log, memory: make(map[element.SymbolElement]language.Bind)}

	return &newCtx
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

func (ctx SystemContext) Get(path element.PathElement) resource.Resource {
	return ctx.root.Get(path)
}

func (ctx SystemContext) Put(path element.PathElement, resource resource.Resource) {
	ctx.root.Put(path, resource)
}
