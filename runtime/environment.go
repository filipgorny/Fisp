package runtime

import "wxl/directives"

type Environment struct {
	Methods  []*directives.Method
	Keywords []*directives.Keyword
}

func NewEnvironment(methods []*directives.Method, keywords []*directives.Keyword) Environment {
	env := Environment{}

	env.Methods = append(env.Methods, methods...)
	env.Keywords = append(env.Keywords, keywords...)

	return env
}

func (e *Environment) Declare(m directives.Method) {
	e.Methods = append(e.Methods, &m)
}
