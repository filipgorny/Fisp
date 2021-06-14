package memory

import (
	"wxl/directives"
	"wxl/element"
	"wxl/language"
)

type NullBind struct {
}

func (bind NullBind) IsElementBind() bool {
	return false
}

func (bind NullBind) IsMethodBind() bool {
	return false
}

func (bind NullBind) IsKeywordBind() bool {
	return false
}

func (bind NullBind) GetElementValue() element.Element {
	return element.NullElement{}
}

func (bind NullBind) GetMethodValue() language.Method {
	return directives.Method{
		Call: func(params []*element.Element, ctx *language.Context) element.Element {
			return element.NullElement{}
		},
	}
}

func (bind NullBind) GetKeywordValue() language.Keyword {
	return directives.Keyword{
		Call: func(params []*element.Element, ctx *language.Context) element.Element {
			return element.NullElement{}
		},
	}
}
