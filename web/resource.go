package web

import (
	"syscall/js"
	"wxl/object"
	"wxl/resource"
)

type WebNodeResource struct {
}

func (r WebNodeResource) Class() object.Class {
	return object.Class{Name: "Web-resource"}
}

func (r WebNodeResource) Get(s resource.Selector) resource.Resource {
	return r
}

func (r WebNodeResource) Put(selector resource.Selector, resource resource.Resource) {
	jsPutMethod.Invoke(js.Null, js.ValueOf(selector), js.ValueOf(resource))
}
