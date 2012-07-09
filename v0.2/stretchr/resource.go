package stretchr

import (
	"fmt"
)

const (

	// IDKey is the key for the ID value of resources.
	IDKey string = "IDKey"

	// EmptyID is the string that represents no ID.
	EmptyID string = ""
)

// Resource represents a single Stretchr resource object.
type Resource struct {

	// data is the internal storage of data.
	data map[string]interface{}

	// path is the path of this resource
	path string

	// session is the session that this resource uses to interact
	// with Stretchr data services.
	session *Session
}

// MakeResource makes a new Resource object with the specified path.
//
// 	r := MakeResource("people")
func MakeResource(session *Session, path string) *Resource {
	r := new(Resource)
	r.path = path
	r.data = make(map[string]interface{})
	r.session = session
	return r
}

// Path gets the path of this resource.
func (r *Resource) Path() string {
	return r.path
}

// Data gets the data for this resource.
func (r *Resource) Data() map[string]interface{} {
	return r.data
}

// Get gets the value of a field from this resource or nil
// if no value was found.
func (r *Resource) Get(key string) interface{} {
	return r.data[key]
}

// Set sets the value of a field for this resource.
//
// Set chains:
//	r.Set("name", "Mat").Set("age", 29)
func (r *Resource) Set(key string, value interface{}) *Resource {
	r.data[key] = value
	return r
}

// Remove deletes a field from this resource.
func (r *Resource) Remove(key string) *Resource {
	delete(r.data, key)
	return r
}

/*
	ID management
*/

// GetID gets the ID for this resource, or returns EmptyID if there isn't one.
func (r *Resource) GetID() string {

	idObj := r.Get(IDKey)

	if idString, ok := idObj.(string); ok {
		return idString
	}

	return EmptyID
}

// SetID sets the ID of this resource.
func (r *Resource) SetID(id string) *Resource {
	return r.Set(IDKey, id)
}

// ClearID clears the internally stored ID for this resource.
func (r *Resource) ClearID() *Resource {
	return r.Remove(IDKey)
}

// HasID gets whether this resource has an ID or not.
func (r *Resource) HasID() bool {
	return r.GetID() != EmptyID
}

/*
	URLs
*/

// AbsoluteURL gets the absolute URL representing this resource.
func (r *Resource) AbsoluteURL() string {

	url := r.session.Url(r.path)

	// add the ID if we have one
	if r.HasID() {
		url = fmt.Sprintf("%s/%s", url, r.GetID())
	}

	return url
}
