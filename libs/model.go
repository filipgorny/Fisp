package libs

import (
	"fmt"
	"wxl/directives"
	"wxl/element"
	"wxl/language"
	"wxl/model"
)

var Model = directives.Method{
	Symbol: element.SymbolElement{Value: "model"},
	Call: func(params []*element.Element, ctx *language.Context) element.Element {

		modelName := *params[1]

		if !modelName.IsSymbol() {
			return Error(ctx, "First argument must be a string.")
		}

		fields := []*model.Field{}

		for i, param := range params {
			if i < 2 {
				continue
			}

			el := *param

			if !el.IsRecord() {
				return Error(ctx, "Invalid model defintion, expected fields records.")
			}

			typeName := el.RecordElementValue().Value().StringValue()

			if !model.ValidTypeName(typeName) {
				return Error(ctx, "Invalid type '"+typeName+"' .")
			}

			field := model.NewField(el.RecordElementValue().Label(), model.TypeByName(typeName))

			fields = append(fields, &field)
		}

		if len(fields) < 1 {
			return Error(ctx, "At least one field must be declared for a model.")
		}

		model := model.NewModel(modelName.StringValue(), fields)

		return element.NewObjectElement(model)
	},
}

var Entity = directives.Method{
	Symbol: element.SymbolElement{Value: "entity"},
	Call: func(params []*element.Element, ctx *language.Context) element.Element {

		modelDef := *params[1]

		if !modelDef.IsObject() {
			fmt.Print(modelDef)
			return Error(ctx, "First argument must be an objet of model.")
		}

		modelDefObj := *modelDef.ObjectElementValue().Object()
		modelObj, ok := modelDefObj.(model.Model)

		if !ok {
			return Error(ctx, "First argument must be a model.")
		}

		var attributes []*model.Attribute

		for i, param := range params {
			if i < 2 {
				continue
			}

			el := *param

			if !el.IsRecord() {
				return Error(ctx, "Invalid entity defintion, expected field record.")
			}

			fieldName := el.RecordElementValue().Label()
			field := modelObj.Field(fieldName)

			if !field.Defined {
				return Error(ctx, "Field "+fieldName+", doesn't exists in the model.")
			}

			value := el.RecordElementValue().Value()

			fmt.Print(value)

			if !field.Type.Validate(value) {
				return Error(ctx, "Invalid value for field "+fieldName+": "+value.StringValue())
			}

			attribute := model.NewAttribute(*field, value)

			attributes = append(attributes, &attribute)
		}

		entity := model.NewEntity(modelObj, attributes)

		fmt.Print(element.NewObjectElement(entity))

		return element.NewObjectElement(entity)
	},
}
