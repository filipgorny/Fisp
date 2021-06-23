package element

type Element interface {
	IsRecord() bool
	IsPath() bool
	IsList() bool
	IsString() bool
	IsNumber() bool
	IsSymbol() bool
	IsError() bool
	IsNull() bool
	IsFunction() bool
	IsObject() bool

	NumberValue() float64
	StringValue() string
	BoolValue() bool

	StringElementValue() *StringElement
	NumberElementValue() *NumberElement
	ListElementValue() *ListElement
	FunctionElementValue() *FunctionElement
	SymbolElementValue() *SymbolElement
	ObjectElementValue() *ObjectElement
	RecordElementValue() *RecordElement
	PathElementValue() *PathElement

	Children() []*Element
}
