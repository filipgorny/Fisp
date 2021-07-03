package web

import (
	"syscall/js"
	"wxl/element"
	"wxl/resource"
	"wxl/serialize"
)

type WebNodeResource struct {
}

func PathTranslate(p element.PathElement) string {
	return p.String()
}

func ResourceTranslate(r resource.Resource) string {
	return serialize.Serialize(r.(element.Element))
}

func (r WebNodeResource) Get(p element.PathElement) resource.Resource {
	return r
}

func (r WebNodeResource) Put(p element.PathElement, resource resource.Resource) {
	js.Global().Call("console.log", ("dupa"))
	js.Global().Call("wxlPutMethod", js.ValueOf(p), js.ValueOf(resource))
}
