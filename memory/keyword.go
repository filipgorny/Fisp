package memory

import (
	"wxl/directives"
	"wxl/element"
	"wxl/language"
)

type KeywordBind struct {
	keyword language.Keyword
}

func NewKeywordBind(keyword language.Keyword) KeywordBind {
	return KeywordBind{
		keyword: keyword,
	}
}

func (bind KeywordBind) IsElementBind() bool {
	return false
}

func (bind KeywordBind) IsKeywordBind() bool {
	return true
}

func (bind KeywordBind) IsMethodBind() bool {
	return false
}

func (bind KeywordBind) GetElementValue() element.Element {
	return element.StringElement{Value: bind.keyword.GetSymbol().StringValue()}
}

func (bind KeywordBind) GetMethodValue() language.Method {
	return directives.Method{
		Call: func(params []*element.Element, ctx *language.Context) element.Element {
			return *params[1]
		},
	}
}

func (bind KeywordBind) GetKeywordValue() language.Keyword {
	return bind.keyword
}
