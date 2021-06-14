package language

import "wxl/element"

type Function interface {
	GetName() element.StringElement
	GetArguments() element.ListElement
	GetBody() element.ListElement
}
