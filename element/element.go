package element

type Element interface {
	IsList() bool
	IsSymbol() bool
	NumberValue() float64
	StringValue() string
	ListElementValue() *ListElement
	Children() []*Element
	SymbolValue() string
}
