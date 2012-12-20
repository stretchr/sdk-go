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

/*
	Session
*/
func (s *Session) Create(resource Resource) (*Response, error) {

	r := s.At(resource.ResourcePath())

	r.httpMethod = HttpMethodPost

	/*
		Set the body
	*/
	var codecErr error
	r.body, codecErr = ObjectToBytes(resource.ResourceData())
	if codecErr != nil {
		return nil, codecErr
	}

	return s.transporter.MakeRequest(r)
}
