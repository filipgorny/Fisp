package web

import "syscall/js"

var jsPutMethod js.Value
var jsGetMethod js.Value

func RegisterPutFunction(fun js.Value) {
	jsPutMethod = fun
}

func RegisterGetFunction(fun js.Value) {
	jsGetMethod = fun
}
