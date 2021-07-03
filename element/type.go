package element

const (
	TYPE_STRING   = "string"
	TYPE_NUMBER   = "number"
	TYPE_BOOL     = "bool"
	TYPE_FUNCTION = "function"
	TYPE_LIST     = "list"
	TYPE_PATH     = "path"
	TYPE_OBJECT   = "object"

	TYPE_DISALLOWED = "disallowed"

	TYPE_UNDEFINED = "undefined"
)

var Types = []string{TYPE_BOOL, TYPE_FUNCTION, TYPE_LIST, TYPE_NUMBER, TYPE_OBJECT, TYPE_PATH, TYPE_STRING}

type TypeElement struct {
	Type string
}

func NewTypeElement(typeName string) TypeElement {
	for _, allowedType := range Types {
		if allowedType == typeName {
			return TypeElement{
				Type: typeName,
			}
		}
	}

	return TypeElement{
		Type: TYPE_DISALLOWED,
	}
}

func (s TypeElement) Name() string {
	return s.Type
}

func (s TypeElement) IsValid() bool {
	for _, allowedType := range Types {
		if allowedType == s.Type {
			return true
		}
	}

	return false
}

func (s TypeElement) Validate(value Element) bool {
	if s.Type == TYPE_BOOL {
		return value.IsBool()
	}

	if s.Type == TYPE_FUNCTION {
		return value.IsFunction()
	}

	if s.Type == TYPE_LIST {
		return value.IsList()
	}

	if s.Type == TYPE_NUMBER {
		return value.IsNumber()
	}

	if s.Type == TYPE_OBJECT {
		return value.IsObject()
	}

	if s.Type == TYPE_PATH {
		return value.IsPath()
	}

	if s.Type == TYPE_STRING {
		return value.IsString()
	}

	if s.Type == TYPE_UNDEFINED {
		return value.IsNull()
	}

	return false
}

func (s TypeElement) IsList() bool {
	return false
}

func (s TypeElement) IsString() bool {
	return false
}

func (s TypeElement) IsSymbol() bool {
	return false
}

func (e TypeElement) IsNumber() bool {
	return false
}

func (s TypeElement) NumberValue() float64 {
	return 0
}

func (s TypeElement) StringValue() string {
	return s.Type
}

func (s TypeElement) SymbolElementValue() *SymbolElement {
	return &SymbolElement{Value: s.Type}
}

func (s TypeElement) IsFunction() bool {
	return false
}

func (s TypeElement) FunctionElementValue() *FunctionElement {
	return &FunctionElement{hasName: false, body: ListElement{}, arguments: ListElement{}}
}

func (s TypeElement) Children() []*Element {
	return nil
}

func (e TypeElement) ListElementValue() *ListElement {
	return nil
}

func (e TypeElement) IsError() bool {
	return false
}

func (e TypeElement) IsNull() bool {
	return false
}

func (e TypeElement) BoolValue() bool {
	return true
}

// object

func (o TypeElement) IsObject() bool {
	return false
}

func (o TypeElement) ObjectElementValue() *ObjectElement {
	return &ObjectElement{}
}

func (e TypeElement) StringElementValue() *StringElement {
	return &StringElement{Value: e.Type}
}

func (e TypeElement) NumberElementValue() *NumberElement {
	return &NumberElement{Value: 1}
}

func (e TypeElement) IsRecord() bool {
	return false
}

func (e TypeElement) RecordElementValue() *RecordElement {
	return &RecordElement{value: e}
}

func (e TypeElement) IsPath() bool {
	return false
}

func (e TypeElement) PathElementValue() *PathElement {
	return &PathElement{}
}

func (e TypeElement) IsType() bool {
	return true
}

func (e TypeElement) TypeElementValue() *TypeElement {
	return &e
}

func (e TypeElement) IsBool() bool {
	return false
}
