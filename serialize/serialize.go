package serialize

import (
	"strings"
	"wxl/element"
)

func Serialize(el element.Element) string {
	json := "{\n"

	if el.IsObject() {
		json += "\tobject: \"" + el.ObjectElementValue().Name.Value + "\",\n"
		json += "\toid: \"" + el.ObjectElementValue().OID + "\",\n"
		json += "\tproperties: [\n"

		for i, prop := range el.ObjectElementValue().Properties {
			json += "\t{\n"
			json += "\t\tname: \"" + prop.Label().Value + "\",\n"

			valJson := ""
			val := prop.Value()

			if val.IsObject() {
				valJson = Serialize(val)
			} else {
				valJson = "\"" + val.StringValue() + "\""
			}

			json += "\t\tvalue: " + valJson + "\n"
			json += "\t}"

			if !(i == len(el.ObjectElementValue().Properties)-1) {
				json += ","
			}

			json += "\n"
		}

		json += "\t]"
	}

	json += "}"

	return json
}

func Flat(json string) string {
	return strings.ReplaceAll((strings.ReplaceAll(json, "\n", "")), "\t", "")
}

func Unserialize(json string) element.Element {
	return element.NewBoolElement(false)
}
