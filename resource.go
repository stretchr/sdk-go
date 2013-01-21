package stretchr

import (
	"github.com/stretchrcom/stew/objects"
	stewstrings "github.com/stretchrcom/stew/strings"
	"github.com/stretchrcom/stretchr-sdk-go/common"
	"strings"
)

// Resource represents a resource in Stretchr.
type Resource struct {
	// Path is the path of this resource.
	path string
	// Data holds the data for the resource.
	data objects.Map
}

// MakeResourceAt makes a new Resource with the specified path.
func MakeResourceAt(path string) *Resource {

	resource := new(Resource)
	resource.path = path
	resource.data = make(map[string]interface{})

	// do we need to set the ID in the data?
	pathSegments := strings.Split(path, common.PathSeparator)
	if len(pathSegments)%2 == 0 {
		// yes -
		resource.data[common.DataFieldID] = pathSegments[len(pathSegments)-1]
	}

	return resource
}

// ResourcePath gets the path for this Resource.
func (r *Resource) ResourcePath() string {
	// TODO: have this cache

	// break the path apart
	pathSegments := strings.Split(r.path, common.PathSeparator)

	// do we have an ID in the data?
	if id, hasId := r.data[common.DataFieldID]; hasId {
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
func (r *Resource) ResourceData() objects.Map {
	return r.data
}

// Get gets a value from the resource.
//
// Keypaths are supported with the dot syntax, for more information see
// http://godoc.org/github.com/stretchrcom/stew/objects#Map.Get
func (r *Resource) Get(keypath string) interface{} {
	return r.data.Get(keypath)
}

// Set sets a value to the specified key and returns the Resource for chaining.
//
// Keypaths are supported with the dot syntax, for more information see
// http://godoc.org/github.com/stretchrcom/stew/objects#Map.Set
func (r *Resource) Set(keypath string, value interface{}) *Resource {
	r.data.Set(keypath, value)
	return r
}

// SetID sets the ID value to the specified string and returns the Resource for chaining.
func (r *Resource) SetID(ID string) *Resource {
	return r.Set(common.DataFieldID, ID)
}
