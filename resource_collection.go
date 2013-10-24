package stretchr

type ResourceCollection struct {
	// Resources represents the resources that make up this collection.
	Resources []*Resource
	// Total gets the total number of resources if known.  To get a total,
	// set total=1 when making the initial request.
	Total float64
}

func NewResourceCollection(resources []*Resource) *ResourceCollection {
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
