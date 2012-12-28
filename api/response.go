package api

import (
	"io/ioutil"
	"net/http"
)

// Repsonse represents a response from a Stretchr request.
type Response struct {
	session      *Session
	httpResponse *http.Response

	// yielded values
	httpStatusCode int
	bodyObject     interface{}
}

func NewResponse(session *Session, httpResponse *http.Response) (*Response, error) {

	response := new(Response)

	// set the session
	response.session = session

	// allow the user to have access to the real underlying HTTP response
	response.httpResponse = httpResponse

	// process the response
	processErr := response.processRequest()

	if processErr != nil {
		return nil, processErr
	}

	return response, nil
}

func (r *Response) processRequest() error {

	// get the repsonse
	bodyBytes, readAll := ioutil.ReadAll(r.httpResponse.Body)

	if readAll != nil {
		return readAll
	}

	var bodyObject interface{}
	unmarshalErr := r.Session().Codec().Unmarshal(bodyBytes, &bodyObject)

	if unmarshalErr != nil {
		return unmarshalErr
	}

	// set the body object
	r.bodyObject = bodyObject

	// all went well
	return nil
}

func (r *Response) HttpResponse() *http.Response {
	return r.httpResponse
}

func (r *Response) Session() *Session {
	return r.session
}

func (r *Response) HttpStatusCode() int {
	return r.httpResponse.StatusCode
}

func (r *Response) BodyObject() interface{} {
	return r.bodyObject
}

func (r *Response) SingleBodyObject() map[string]interface{} {
	return r.BodyObject().(map[string]interface{})
}

func (r *Response) MultipleBodyObjects() []map[string]interface{} {

	bodyObjs := r.BodyObject().([]interface{})
	typedBodyObjs := make([]map[string]interface{}, len(bodyObjs))
	for objIndex, obj := range bodyObjs {
		typedBodyObjs[objIndex] = obj.(map[string]interface{})
	}

	return typedBodyObjs
}
