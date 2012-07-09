package stretchr

import (
	"errors"
	"fmt"
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

// AbsoluteURL gets the absolute URL representing this resource.
func (r *Resource) AbsoluteURL() string {
	return r.session.Url(r.path)
}

/*
	Action methods
*/

// Create creates a new resource.
func (r *Resource) Create() error {

	json, jsonErr := toJson(r.data)

	if jsonErr != nil {
		return jsonErr
	}

	response, _, requestErr := ActiveRequester.MakeRequest(CreateMethod, r.AbsoluteURL(), json)

	if requestErr != nil {
		return requestErr
	}

	if response.Worked {

		// get the new ID
		r.Set("~id", response.Data["~id"])

	} else {

		if len(response.Errors) > 0 {
			return errors.New(fmt.Sprintf("%s", response.Errors[0].(map[string]interface{})["Message"]))
		} else {
			return errors.New("Unknown error")
		}

	}

	// all OK
	return nil

}
