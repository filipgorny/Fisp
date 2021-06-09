package language

import "wxl/core"

type BuiltInMethodCallResult struct {
	Element core.Element
	Error   string
}

type BuiltInMethod struct {
	Symbol core.SymbolElement
	Call   func(params []*core.Element) BuiltInMethodCallResult
}

type Spec struct {
	Definitions []BuiltInMethod
}

func CreateSpec(methods []BuiltInMethod) Spec {
	var spec = Spec{}

	for _, method := range methods {
		spec.Definitions = append(spec.Definitions, method)
	}

	return spec
}
