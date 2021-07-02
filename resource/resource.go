package resource

type Resource interface {
}

type Container interface {
	Resource

	Get(selector Selector) Resource
	Put(selector Selector, resource Resource)
}
