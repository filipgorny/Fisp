package libs

import (
	"wxl/directives"
	"wxl/element"
	"wxl/language"
)

var Field = directives.Method{
	Symbol: element.SymbolElement{Value: "field"},
	Call: func(params []*element.Element, ctx *language.Context) element.Element {
		el := *params[1]

		if !(el.IsRecord() || el.RecordElementValue().Value().TypeElementValue().IsValid()) {
			return element.NewErrorElement("Field argument must be a record with a type as value.")
		}

		fieldName := el.RecordElementValue().Label()
		fieldType := el.RecordElementValue().Value().TypeElementValue()

		if !fieldType.IsValid() {
			return element.NewErrorElement("Invalid type '" + fieldType.Name() + "' .")
		}

		field := element.NewObjectElement("field", []element.RecordElement{element.NewRecordElement("name", fieldName), element.NewRecordElement("type", fieldType)})

		return field
	},
}

var Model = directives.Method{
	Symbol: element.SymbolElement{Value: "model"},
	Call: func(params []*element.Element, ctx *language.Context) element.Element {

		modelName := *params[1]

		if !modelName.IsSymbol() {
			return element.NewErrorElement("First argument must be a string.")
		}

		fields := element.NewListElement()

		for i, param := range params {
			if i < 2 {
				continue
			}

			el := *param

			if !el.ObjectElementValue().IsA("field") {
				return element.NewErrorElement("Invalid model defintion, expected fields.")
			}

			field := el.ObjectElementValue()

			fields.Push(field)
		}

		if fields.Length() < 1 {
			return element.NewErrorElement("At least one field must be declared for a model.")
		}

		model := element.NewObjectElement("model", []element.RecordElement{element.NewRecordElement("name", modelName), element.NewRecordElement("fields", fields)})

		return model
	},
}

var Entity = directives.Method{
	Symbol: element.SymbolElement{Value: "entity"},
	Call: func(params []*element.Element, ctx *language.Context) element.Element {
		// entity .task label: test done: false
		modelDef := *params[1]

		if !modelDef.ObjectElementValue().IsA("model") {
			return element.NewErrorElement("First argument must be a model.")
		}

		var model element.ObjectElement = *modelDef.ObjectElementValue()

		var modelFields element.Element = model.Get("fields")
		modelFieldsList := modelFields.ListElementValue()

		if modelFieldsList.Length() == 0 {
			return element.NewErrorElement("Invalid fields list from the model.")
		}

		validField := func(name string, value element.Element) bool {
			for _, element := range modelFieldsList.All() {
				field := *element

				if field.ObjectElementValue().Get("name").StringValue() == name {
					if field.ObjectElementValue().Get("type").TypeElementValue().Validate(value) {
						return true
					}
				}
			}

			return false
		}

		attributes := element.NewListElement()

		for i, param := range params {
			if i < 2 {
				continue
			}

			el := *param

			if !el.IsRecord() {
				return element.NewErrorElement("Invalid entity defintion, expected field record.")
			}

			fieldName := el.RecordElementValue().Label()
			fieldValue := el.RecordElementValue().Value()

			if !validField(fieldName.Value, fieldValue) {
				return element.NewErrorElement("Field ", fieldName.Value, ", doesn't exists in the model.")
			}

			attribute := element.NewObjectElement("attribute", []element.RecordElement{element.NewRecordElement("name", fieldName), element.NewRecordElement("value", fieldValue)})

			attributes.Push(attribute)
		}

		// getterCode := `fun (this, name)
		//   (each this.attributes
		//     (== .name name
		//       (return .))
		//     )

		// `

		// getter := element.NewFunctionElement(true, "*get", core.Parse(core.Tokenize(string(getterCode))).Elements)

		return element.NewObjectElement("entity", []element.RecordElement{element.NewRecordElement("model", modelDef), element.NewRecordElement("attributes", attributes)})
	},
}
