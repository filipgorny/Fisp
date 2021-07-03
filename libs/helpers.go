package libs

import (
	"wxl/core"
	"wxl/directives"
	"wxl/element"
	"wxl/execution"
	"wxl/runtime"
)

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

		&Field,
		&Model,
		&Entity,

		&Put,
		&Get,
	}, []*directives.Keyword{
		&Fun,
	})

	code := core.Parse(core.Tokenize(string(s)))

	return execution.Run(env, code)
}
