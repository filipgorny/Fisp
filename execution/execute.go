package execution

import (
	"wxl/core"
	"wxl/element"
	"wxl/logging"
	"wxl/runtime"
	"wxl/stack"
)

func Run(env runtime.Environment, codeTree core.ExprList) element.Element {
	var runtimeLog = logging.Log{}

	resultList := element.ListElement{}
	ctx := stack.NewContext(env, &runtimeLog)

	for _, element := range codeTree.Elements {
		if element.IsList() {
			newCtx := *ctx.Branch()
			result := runtime.Evaluate(element.ListElementValue(), newCtx)

			resultList.Push(result)
		}
	}

	resultElement := *resultList.Last()

	runtimeLog.Result(resultElement.StringValue())

	runtimeLog.Flush()

	return resultElement
}
