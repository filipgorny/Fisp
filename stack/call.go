package stack

import (
	"wxl/element"
)

type Call func(params []*element.Element, ctx *Context) element.Element
