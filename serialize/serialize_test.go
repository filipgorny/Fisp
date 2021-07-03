package serialize

import (
	"testing"
	"wxl/element"
)

func TestSerialize(t *testing.T) {
	object := element.NewObjectElement("test-object", nil)

	object.Set("property1", element.NewStringElement("BLABLABLA"))
	object.Set("property2", element.NewBoolElement(true))

	subobject := element.NewObjectElement("subobject", nil)
	subobject.Set("property1", element.NewStringElement("LOLOLO"))

	object.Set("property3", subobject)

	json := Serialize(object)

	t.Log(Flat(json))

	return
}
