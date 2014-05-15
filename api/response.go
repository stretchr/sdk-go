package api

import (
	"io/ioutil"
	"net/http"
)

// Response represents a response from a Stretchr request.
type Response struct {
	session      *Session
	httpResponse *http.Response

	// yielded values
	httpStatusCode int
	bodyObject     ResponseObject
}

// NewResponse creates a new Response object from the given Session and http.Response.
func NewResponse(session *Session, httpResponse *http.Response) (*Response, error) {

	response := new(Response)

	// set the session
	response.session = session

	// allow the user to have access to the real underlying HTTP response
	response.httpResponse = httpResponse

	// process the response
	processErr := response.processResponse()

	if processErr != nil {
		return nil, processErr
	}

	return response, nil
}

// processRequest processes the http.Response and builds the current Response object.
func (r *Response) processResponse() error {

	// get the repsonse
	bodyBytes, readAll := ioutil.ReadAll(r.httpResponse.Body)

	if readAll != nil {
		return readAll
	}

	var bodyObject map[string]interface{}
	unmarshalErr := r.Session().Codec().Unmarshal(bodyBytes, &bodyObject)

	if unmarshalErr != nil {
		return unmarshalErr
	}

	// set the body object
	r.bodyObject = ResponseObject(bodyObject)

	// all went well
	return nil
}

// HTTPResponse gets the http.Response that this Response represents.
func (r *Response) HTTPResponse() *http.Response {
	return r.httpResponse
}

// Session gets the Session that was responsible for obtaining this Response.
func (r *Response) Session() *Session {
	return r.session
}

// HTTPStatusCode gets the HTTP Status Code of this response.
func (r *Response) HTTPStatusCode() int {
	return r.httpResponse.StatusCode
}

// BodyObject gets a real object unmarshalled from the response data.
func (r *Response) BodyObject() ResponseObject {
	return r.bodyObject
}
