package object

type Null struct {
}

func (n Null) Class() Class {
	return Class{Name: "null"}
}
