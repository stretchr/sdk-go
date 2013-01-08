package stretchr

type ResourceCollection struct {
	Resources []*Resource
}

func MakeResourceCollection(resources []*Resource) *ResourceCollection {
	c := new(ResourceCollection)
	c.Resources = resources
	return c
}
