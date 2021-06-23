package helpers

import (
	"wxl/element"
	"wxl/language"
)

func Error(ctx *language.Context, message string) element.ErrorElement {
	context := *ctx

	context.Error(message)

	return element.ErrorElement{Message: message}
}
