package api

import (
	"github.com/stretchrcom/stretchr-sdk-go/common"
)

// Read executes the Request with a GET method, and returns the Response, or an error
// if something went wrong communicating with Stretchr.
func (r *Request) Read() (*Response, error) {

	// set the HTTP method
	r.httpMethod = common.HttpMethodGet

	// get the transporter to do the work
	return r.session.transporter.MakeRequest(r)
}

// Delete executes the Request with a DELETE method, and returns the Response, or an error
// if something went wrong communicating with Stretchr.
func (r *Request) Delete() (*Response, error) {

	// set the HTTP method
	r.httpMethod = common.HttpMethodDelete

	// get the transporter to do the work
	return r.session.transporter.MakeRequest(r)
}

// CreateMany executes the Request with a POST method, and returns the Response, or an error
// if something went wrong communicating with Stretchr.
//
// The ResourcePath is taken from the Request, and the paths of individual resources are
// ignored.
func (r *Request) CreateMany(resources []Resource) (*Response, error) {

	// set the HTTP method
	r.httpMethod = common.HttpMethodPost

	// collect the data objects
	var dataObjects []interface{} = make([]interface{}, len(resources))
	for resourceIndex, resource := range resources {
		dataObjects[resourceIndex] = resource.ResourceData()
	}
	r.setBodyObject(dataObjects)

	// get the transporter to do the work
	return r.session.transporter.MakeRequest(r)
}

// createOrReplace performs the Create or Replace action (they are currently the same)
func (r *Request) createOrReplace(resource Resource) (*Response, error) {

	r.httpMethod = common.HttpMethodPost
	r.setBodyObject(resource.ResourceData())

	return r.session.transporter.MakeRequest(r)
}

// Create tells Stretchr to create the specified resource.
func (r *Request) Create(resource Resource) (*Response, error) {
	return r.createOrReplace(resource)
}

// Replace tells Stretchr to replace the specified resource.
func (r *Request) Replace(resource Resource) (*Response, error) {
	return r.createOrReplace(resource)
}

// Update tells Stretchr to update the specified resource.
func (r *Request) Update(resource Resource) (*Response, error) {

	r.httpMethod = common.HttpMethodPut
	r.setBodyObject(resource.ResourceData())

	return r.session.transporter.MakeRequest(r)
}
