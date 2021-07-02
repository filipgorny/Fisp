package object

import "wxl/serialization"

type Class struct {
	Name string
}

type Object interface {
	serialization.Serializable

	Class() Class
}
