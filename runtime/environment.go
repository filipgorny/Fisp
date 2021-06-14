package runtime

type Environment struct {
	Methods []*Method
}

func NewEnvironment(methods []*Method) Environment {
	env := Environment{}

	env.Methods = append(env.Methods, methods...)

	return env
}

func (e *Environment) Declare(m Method) {
	e.Methods = append(e.Methods, &m)
}
