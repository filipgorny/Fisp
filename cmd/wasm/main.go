package main

import (
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

func wxlConstructor() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		this.Set("registerPutMethod", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			web.RegisterPutFunction(args[0])

			return js.Null()
		}))

		this.Set("registerGetMethod", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			web.RegisterGetFunction(args[0])

			return js.Null()
		}))

		return this
	})

	return jsonFunc
}

func main() {
	done := make(chan struct{}, 0)
	defer close(done)

	js.Global().Set("wxlSetPutMethod", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		web.RegisterPutFunction(args[0])

		return js.Null()
	}))

	js.Global().Set("wxlSetGetMethod", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		web.RegisterGetFunction(args[0])

		return js.Null()
	}))

	js.Global().Set("runWxl", createWxlRunnerFunction())

	<-done
}
