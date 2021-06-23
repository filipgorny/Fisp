// package resource

// import (
// 	"wxl/directives"
// 	"wxl/element"
// 	"wxl/language"
// 	"wxl/libs/helpers"
// 	"wxl/resource"
// )

// var Get = directives.Method{
// 	Symbol: element.SymbolElement{Value: "get"},
// 	Call: func(params []*element.Element, ctx *language.Context) element.Element {
// 		currentContext := *ctx
// 		firstArgument := *params[1]
// 		secondArgument := *params[2]

// 		if !firstArgument.IsPath() {
// 			return helpers.Error(ctx, "First argument must be a path.")
// 		}

// 		var subject resource.Resource

// 		if len(params) < 3 {
// 			subject = currentContext
// 		} else {
// 			objectEl := secondArgument.ObjectElementValue()

// 			if objectEl.Object() == nil {
// 				return element.NullElement{}
// 			}

// 			object := *objectEl.Object()

// 			subject = object.(resource.Resource)

// 			if subject == nil {
// 				return helpers.Error(ctx, "Invalid subject.")
// 			}
// 		}

// 		result := subject.Get(firstArgument.PathElementValue().Path())

// 		return element.NewObjectElement(result)
// 	},
// }

// // var Put = directives.Method{
// // 	Symbol: element.SymbolElement{Value: "put"},
// // 	Call: func(params []*element.Element, ctx *language.Context) element.Element {
// // 		currentContext := *ctx
// // 		firstArgument := *params[1]
// // 		secondArgument := *params[2]

// // 		if !firstArgument.IsSymbol() {
// // 			return helpers.Error(ctx, "First argument must be a symbol.")
// // 		}

// // 		if len(params) < 3 {
// // 			return helpers.Error(ctx, "Not enough arguments.")
// // 		}

// // 		object := *secondArgument.ObjectElementValue().Object()

// // 		currentContext.Declare(*firstArgument.SymbolElementValue(), memory.NewObjectBind(&object))

// // 		return secondArgument
// // 	},
// // }
