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

	ctx.Log("start")

	for _, element := range codeTree.Elements {
		ctx.Log("first element")
		ctx.Log(element)
		if element.IsList() {
			ctx.Log("element is list: " + element.StringValue())
			newCtx := ctx.Branch()
			result := runtime.Evaluate(element.ListElementValue(), newCtx)

			ctx.Log("evaluation result: " + result.StringValue())

			resultList.Push(result)
		}
	}

	resultElement := *resultList.Last()

	runtimeLog.Result(resultElement.StringValue())

	runtimeLog.Flush()

	return resultElement
}
