package serialization

type Serializable interface {
	Serialize() SerializedData
	UnSerialize(data SerializedData) Serializable
}

type SerializedData struct {
	Attributes []DataAttribute
}

type DataAttribute struct {
	Name  string
	Value string
}
