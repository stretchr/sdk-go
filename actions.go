package stretchr

/*
	Developers notice
	=================
	This file contains actions like Read, Create, Update, Replace and Delete.  They happen to
	appear on different objects, but are in one place here for simplicity's sake.
*/

/*
	Request
*/

// Read executes the Request with a GET method, and returns the Response, or an error
// if something went wrong communicating with Stretchr.
func (r *Request) Read() (*Response, error) {

	// set the HTTP method
	r.httpMethod = HttpMethodGet

	// get the transporter to do the work
	return r.session.transporter.MakeRequest(r)
}

// Delete executes the Request with a DELETE method, and returns the Response, or an error
// if something went wrong communicating with Stretchr.
func (r *Request) Delete() (*Response, error) {

	// set the HTTP method
	r.httpMethod = HttpMethodDelete

	// get the transporter to do the work
	return r.session.transporter.MakeRequest(r)
}

// CreateMany executes the Request with a POST method, and returns the Response, or an error
// if something went wrong communicating with Stretchr.
func (r *Request) CreateMany(resources []Resource) (*Response, error) {

	// set the HTTP method
	r.httpMethod = HttpMethodPost

	// collect the data objects
	var dataObjects []interface{} = make([]interface{}, len(resources))
	for resourceIndex, resource := range resources {
		dataObjects[resourceIndex] = resource.ResourceData()
	}
	r.setBodyObject(dataObjects)

	// get the transporter to do the work
	return r.session.transporter.MakeRequest(r)
}

/*
	Session
*/

// createOrReplace performs the Create or Replace action (they are currently the same)
func (s *Session) createOrReplace(resource Resource) (*Response, error) {

	r := s.At(resource.ResourcePath())

	r.httpMethod = HttpMethodPost
	r.setBodyObject(resource.ResourceData())

	return s.transporter.MakeRequest(r)
}

// Create tells Stretchr to create the specified resource.
func (s *Session) Create(resource Resource) (*Response, error) {
	return s.createOrReplace(resource)
}

// Replace tells Stretchr to replace the specified resource.
func (s *Session) Replace(resource Resource) (*Response, error) {
	return s.createOrReplace(resource)
}

// Update tells Stretchr to update the specified resource.
func (s *Session) Update(resource Resource) (*Response, error) {

	r := s.At(resource.ResourcePath())

	r.httpMethod = HttpMethodPut
	r.setBodyObject(resource.ResourceData())

	return s.transporter.MakeRequest(r)
}
