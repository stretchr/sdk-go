package stretchr

type ResourceCollection struct {
	Resources []*Resource
}

func MakeResourceCollection(resources []*Resource) *ResourceCollection {
	c := new(ResourceCollection)
	if resources == nil {
		resources = make([]*Resource, 0)
	}
	c.Resources = resources
	return c
}

func (r *ResourceCollection) AddResource(resource *Resource) {
	r.Resources = append(r.Resources, resource)
}

func (r *ResourceCollection) IsEmpty() bool {
	return len(r.Resources) == 0
}
