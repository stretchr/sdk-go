package api

import (
	"github.com/stretchr/sdk-go/common"
)

// Read executes the Request with a GET method, and returns the Response, or an error
// if something went wrong communicating with Stretchr.
func (r *Request) Read() (*Response, error) {

	// set the HTTP method
	r.httpMethod = common.HTTPMethodGet

	// get the transporter to do the work
	return r.session.transporter.MakeRequest(r)
}

// Delete executes the Request with a DELETE method, and returns the Response, or an error
// if something went wrong communicating with Stretchr.
func (r *Request) Delete() (*Response, error) {

	// set the HTTP method
	r.httpMethod = common.HTTPMethodDelete

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
	r.httpMethod = common.HTTPMethodPost

	// collect the data objects
	dataObjects := make([]interface{}, len(resources))
	for resourceIndex, resource := range resources {
		dataObjects[resourceIndex] = resource.ResourceData()
	}
	r.setBodyObject(dataObjects)

	// get the transporter to do the work
	return r.session.transporter.MakeRequest(r)
}

// Create tells Stretchr to create the specified resource.
// If an ID is provided, a resource with that ID will be created.
// If the ID exists in the database, the creation will fail.
func (r *Request) Create(resource Resource) (*Response, error) {
	r.httpMethod = common.HTTPMethodPost
	r.setBodyObject(resource.ResourceData())

	return r.session.transporter.MakeRequest(r)
}

// Replace tells Stretchr to replace the specified resource.
// If the resource does not exist, it will be created.
func (r *Request) Replace(resource Resource) (*Response, error) {
	r.httpMethod = common.HTTPMethodPut
	r.setBodyObject(resource.ResourceData())

	return r.session.transporter.MakeRequest(r)
}

// Update tells Stretchr to update the specified resource.
// An ID is required. If the ID doesn't not exist in the database,
// this call is an error.
func (r *Request) Update(resource Resource) (*Response, error) {
	r.httpMethod = common.HTTPMethodPatch
	r.setBodyObject(resource.ResourceData())

	return r.session.transporter.MakeRequest(r)
}
