package stretchr

import (
	"strings"
)

// ResourceObject represents a general Resource implementation.
type ResourceObject struct {
	// Path is the path of this resource.
	Path string
	// Data holds the data for the resource.
	Data map[string]interface{}
}

// MakeResourceAt makes a new Resource with the specified path.
func MakeResourceAt(path string) *ResourceObject {
	resource := new(ResourceObject)
	resource.Path = path
	resource.Data = make(map[string]interface{})

	// do we need to set the ID in the data?
	pathSegments := strings.Split(path, pathSeparator)
	if len(pathSegments)%2 == 0 {
		// yes -
		resource.Data[dataFieldID] = pathSegments[len(pathSegments)-1]
	}

	return resource
}

// ResourcePath gets the path for this Resource.
func (r *ResourceObject) ResourcePath() string {

	// break the path apart
	pathSegments := strings.Split(r.Path, pathSeparator)

	// do we have an ID in the data?
	if id, hasId := r.Data[dataFieldID]; hasId {
		// do we have an ID in the path?
		if len(pathSegments)%2 == 0 {
			// update the ID
			pathSegments[len(pathSegments)-1] = id.(string)
		} else {
			// add the ID
			pathSegments = append(pathSegments, id.(string))
		}
	}

	return JoinStrings(pathSeparator, pathSegments...)
}

// ResourceData gets the data for this Resource.
func (r *ResourceObject) ResourceData() map[string]interface{} {
	return r.Data
}
