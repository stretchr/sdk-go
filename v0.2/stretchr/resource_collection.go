package stretchr

type ResourceCollection struct {
	Resources []*Resource
}

func MakeResourceCollection() *ResourceCollection {
	return new(ResourceCollection)
}
