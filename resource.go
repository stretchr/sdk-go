package stretchr

import (
	stewstrings "github.com/stretchrcom/stew/strings"
	"github.com/stretchrcom/stretchr-sdk-go/common"
	"strings"
)

// Resource represents a resource in Stretchr.
type Resource struct {
	// Path is the path of this resource.
	Path string
	// Data holds the data for the resource.
	Data map[string]interface{}
}

// MakeResourceAt makes a new Resource with the specified path.
func MakeResourceAt(path string) *Resource {

	resource := new(Resource)
	resource.Path = path
	resource.Data = make(map[string]interface{})

	// do we need to set the ID in the data?
	pathSegments := strings.Split(path, common.PathSeparator)
	if len(pathSegments)%2 == 0 {
		// yes -
		resource.Data[common.DataFieldID] = pathSegments[len(pathSegments)-1]
	}

	return resource
}

// ResourcePath gets the path for this Resource.
func (r *Resource) ResourcePath() string {

	// break the path apart
	pathSegments := strings.Split(r.Path, common.PathSeparator)

	// do we have an ID in the data?
	if id, hasId := r.Data[common.DataFieldID]; hasId {
		// do we have an ID in the path?
		if len(pathSegments)%2 == 0 {
			// update the ID
			pathSegments[len(pathSegments)-1] = id.(string)
		} else {
			// add the ID
			pathSegments = append(pathSegments, id.(string))
		}
	}

	return stewstrings.JoinStrings(common.PathSeparator, pathSegments...)
}

// ResourceData gets the data for this Resource.
func (r *Resource) ResourceData() map[string]interface{} {
	return r.Data
}
