package resource

type PathSelector struct {
	elements []string
}

func NewPath(elements []string) PathSelector {
	var pathElements []string

	for i, el := range elements {
		if el == "" && i == 0 {
			continue
		}

		if el == "" {
			continue
		}

		pathElements = append(pathElements, el)
	}

	return PathSelector{elements: pathElements}
}

func (p PathSelector) Elements() []string {
	return p.elements
}

func (p PathSelector) String() string {
	var path string

	for _, el := range p.elements {
		path += "."
		path += el
	}

	return path
}
