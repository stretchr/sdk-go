package stretchr

import (
	"fmt"
	"github.com/stretchrcom/stew/objects"
	stewstrings "github.com/stretchrcom/stew/strings"
	"github.com/stretchrcom/stretchr-sdk-go/api"
	"github.com/stretchrcom/stretchr-sdk-go/common"
	"strings"
	"time"
)

const (
	// NoID represents the string indicating the resource has no ID.
	NoID string = ""
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

// GetString gets a strongly typed value from the data of this Resource
// or panics if that is impossible.
func (r *Resource) GetString(keypath string) string {
	if val, ok := r.Get(keypath).(string); ok {
		return val
	}
	panic(fmt.Sprintf("stretchr: Cannot GetString on %s.", r.Get(keypath)))
}

// GetNumber gets a strongly typed value from the data of this Resource
// or panics if that is impossible.
func (r *Resource) GetNumber(keypath string) float64 {
	if val, ok := r.Get(keypath).(float64); ok {
		return val
	}
	panic(fmt.Sprintf("stretchr: Cannot GetNumber on %s.", r.Get(keypath)))
}

// GetBool gets a strongly typed value from the data of this Resource
// or panics if that is impossible.
func (r *Resource) GetBool(keypath string) bool {
	if val, ok := r.Get(keypath).(bool); ok {
		return val
	}
	panic(fmt.Sprintf("stretchr: Cannot GetBool on %s.", r.Get(keypath)))
}

// GetTime gets a strongly typed value from the data of this Resource
// or panics if that is impossible.
func (r *Resource) GetTime(keypath string) time.Time {

	if floatTime, floatTimeOk := r.Get(keypath).(float64); floatTimeOk {
		return time.Unix(int64(floatTime), 0)
	}

	panic(fmt.Sprintf("stretchr: Cannot GetTime on %s.", r.Get(keypath)))
}

// ID gets the ID string for this resource.
//
// If this resource doesn't have an ID (which can happen if it has not yet
// been persisted) then `NoID` will be returned.
func (r *Resource) ID() string {
	if id, ok := r.data[common.DataFieldID]; ok {
		if idString, typeOk := id.(string); typeOk {
			return idString
		}
	}
	// no ID
	return NoID
}

// Set sets a value to the specified key and returns the Resource for chaining.
//
// Keypaths are supported with the dot syntax, for more information see
// http://godoc.org/github.com/stretchrcom/stew/objects#Map.Set
func (r *Resource) Set(keypath string, value interface{}) *Resource {

	toStore := value

	switch value.(type) {
	case int:
		toStore = float64(value.(int))
	case int8:
		toStore = float64(value.(int8))
	case int16:
		toStore = float64(value.(int16))
	case int32:
		toStore = float64(value.(int32))
	case int64:
		toStore = float64(value.(int64))
	case uint:
		toStore = float64(value.(uint))
	case uint8:
		toStore = float64(value.(uint8))
	case uint16:
		toStore = float64(value.(uint16))
	case uint32:
		toStore = float64(value.(uint32))
	case uint64:
		toStore = float64(value.(uint64))
	case float32:
		toStore = float64(value.(float32))
	}

	r.data.Set(keypath, toStore)
	return r
}

// SetID sets the ID value to the specified string and returns the Resource for chaining.
func (r *Resource) SetID(ID string) api.Resource {
	return r.Set(common.DataFieldID, ID)
}
