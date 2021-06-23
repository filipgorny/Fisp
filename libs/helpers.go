package libs

import (
	"wxl/core"
	"wxl/directives"
	"wxl/element"
	"wxl/execution"
	"wxl/language"
	"wxl/runtime"
)

func Error(ctx *language.Context, message string) element.ErrorElement {
	context := *ctx

	context.Error(message)

	return element.ErrorElement{Message: message}
}

func runCode(s string) element.Element {
	env := runtime.NewEnvironment([]*directives.Method{
		&Add,
		&Substract,
		&Divide,
		&Multiple,

		&Equal,
		&NotEqual,
		&LeftBigger,
		&RightBigger,
		&IfTrue,

		&Declare,

		&Model,
		&Entity,
	}, []*directives.Keyword{
		&Fun,
	})

	code := core.Parse(core.Tokenize(string(s)))

	return execution.Run(env, code)
}
