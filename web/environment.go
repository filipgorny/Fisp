package web

import (
	"wxl/directives"
	"wxl/libs"
	"wxl/runtime"
)

var Environment = runtime.NewEnvironment(
	[]*directives.Method{
		&libs.Add,
		&libs.Substract,
		&libs.Divide,
		&libs.Multiple,

		&libs.Equal,
		&libs.NotEqual,
		&libs.LeftBigger,
		&libs.RightBigger,
		&libs.IfTrue,

		&libs.Declare,

		&libs.Model,
		&libs.Entity,

		&libs.Put,
		&libs.Get,
	},
	[]*directives.Keyword{
		&libs.Fun,
	},
)
