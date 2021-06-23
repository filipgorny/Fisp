package model

import (
	"fmt"
	"wxl/object"
	"wxl/resource"

	uuid "github.com/nu7hatch/gouuid"
)

type Id = string

type Entity struct {
	object.Object
	resource.Resource

	Id         Id
	Model      Model
	Attributes []*Attribute
}

func NewEntity(model Model, attributes []*Attribute) Entity {
	uuid, _ := uuid.NewV4()

	return Entity{
		Id:         uuid.String(),
		Model:      model,
		Attributes: attributes,
	}
}

func (e Entity) Class() object.Class {
	return object.Class{Name: "entity"}
}

func (e Entity) Attribute(name string) *Attribute {
	for _, a := range e.Attributes {
		if a.Field.name == name {
			return a
		}
	}

	return &UndefinedAttribute
}

func (e Entity) Get(s resource.Selector) resource.Resource {
	fmt.Println("Entity get")
	fmt.Print(s.String())

	// if reflect.TypeOf(s).Name() == "PathSelector" {
	// 	path := s.(resource.PathSelector)

	// 	field := path.Elements()[0]

	// 	return resource.NewElementResource(e.Attribute(field).Value)
	// }

	return resource.NullResource{}
}
