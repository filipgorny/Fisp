package object

type Class struct {
	Name string
}

type Object interface {
	Class() Class
}
