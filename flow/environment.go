package flow

type Method struct {
	name string
	Call Call
}

type Environment struct {
	Methods []Method
}

func (e Environment) HasDefined(name string) bool {
	for _, def := range e.Methods {
		if def.name == name {
			return true
		}
	}

	return false
}
