package flow

import "rpl/element"

type Call func(params []*element.Element, ctx *Context) element.Element
