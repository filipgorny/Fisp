package resource

import "wxl/object"

type NullResource struct {
	object.Null
}

func (n NullResource) Class() object.Class {
	return object.Class{Name: "null-resource"}
}

func (n NullResource) Get(s Selector) Resource {
	return NullResource{}
}
