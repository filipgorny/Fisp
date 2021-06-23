package libs

// import (
// 	"testing"
// 	"wxl/core"
// 	"wxl/directives"
// 	"wxl/element"
// 	"wxl/execution"
// 	"wxl/runtime"
// )

// func runCode(s string) element.Element {
// 	env := runtime.NewEnvironment([]*directives.Method{
// 		&Get,
// 	}, []*directives.Keyword{})

// 	code := core.Parse(core.Tokenize(string(s)))

// 	return execution.Run(env, code)
// }

// func TestFun(t *testing.T) {
// 	// result := runCode(`(
// 	// 	(set task (model label: string done: bool)))
// 	// 	(new task label: "test" done: false)
// 	// )`)

// 	// if result.IsError() || result.NumberValue() != 60 {
// 	// 	t.Error(result.StringValue())
// 	// }
// }
