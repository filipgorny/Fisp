package element

type Element interface {
	IsList() bool
	IsSymbol() bool
	IsError() bool
	IsNull() bool
	NumberValue() float64
	StringValue() string
	ListElementValue() *ListElement
	BoolValue() bool
	Children() []*Element
	SymbolElementValue() *SymbolElement
}
