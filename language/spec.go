package language

import (
	"rpl/element"
)

type RplCallResult struct {
	Element element.Element
	Success bool
}

type RplMethod struct {
	Symbol element.SymbolElement
	Call   func(params []*element.Element) RplCallResult
}

type Spec struct {
	Definitions []RplMethod
}

func CreateSpec(methods []RplMethod) Spec {
	var spec = Spec{}

	for _, method := range methods {
		spec.Definitions = append(spec.Definitions, method)
	}

	return spec
}
