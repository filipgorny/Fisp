package element

type Element interface {
	IsList() bool
	IsSymbol() bool
	IsError() bool
	IsNull() bool
	IsFunction() bool
	NumberValue() float64
	StringValue() string
	ListElementValue() *ListElement
	BoolValue() bool
	FunctionElementValue() *FunctionElement
	Children() []*Element
	SymbolElementValue() *SymbolElement
}
