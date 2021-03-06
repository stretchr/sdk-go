package api

import (
	"strings"

	"github.com/stretchr/objx"
	"github.com/stretchr/sdk-go/common"
	stewstrings "github.com/stretchr/stew/strings"
)

// TestResource is an object used solely for testing the Resource objects
type TestResource struct {
	// Path is the path of this resource.
	Path string
	// Data holds the data for the resource.
	Data objx.Map
}

// MakeTestResourceAt makes a new test resource
func MakeTestResourceAt(path string) *TestResource {
	resource := new(TestResource)
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
func (r *TestResource) ResourcePath() string {

	// break the path apart
	pathSegments := strings.Split(r.Path, common.PathSeparator)

	// do we have an ID in the data?
	if id, hasID := r.Data[common.DataFieldID]; hasID {
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
func (r *TestResource) ResourceData() objx.Map {
	return r.Data
}

// ID Gets the ID for this resource
func (r *TestResource) ID() string {
	return r.Data[common.DataFieldID].(string)
}

// SetID sets the ID for this resource
func (r *TestResource) SetID(ID string) Resource {
	r.Data[common.DataFieldID] = ID
	return r
}
