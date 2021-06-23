package resource

type Resource interface {
}

type Container interface {
	Resource

	Get(selector Selector) Resource
}
