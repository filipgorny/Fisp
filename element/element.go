package element

type Element interface {
	IsList() bool
	IsSymbol() bool
	IsError() bool
	NumberValue() float64
	StringValue() string
	ListElementValue() *ListElement
	BoolValue() bool
	Children() []*Element
	SymbolValue() string
}
