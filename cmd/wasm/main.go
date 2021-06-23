package main

import (
	"fmt"
	"syscall/js"
	"wxl/core"
	"wxl/execution"
	"wxl/web"
)

func createWxlRunnerFunction() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 1 {
			return "Invalid no of arguments passed"
		}

		inputCode := args[0].String()

		result := execution.Run(web.Environment, core.Parse(core.Tokenize(inputCode)))

		return result.StringValue()
	})

	return jsonFunc
}

func main() {
	fmt.Println("Wxl Web Assembly")
	js.Global().Set("runWxl", createWxlRunnerFunction())
	select {}
}
