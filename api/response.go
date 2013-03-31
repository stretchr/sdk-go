package api

import (
	"errors"
	"fmt"
	"github.com/stretchrcom/sdk-go/common"
	"github.com/stretchrcom/signature"
	"io/ioutil"
	"net/http"
)

// Repsonse represents a response from a Stretchr request.
type Response struct {
	session      *Session
	httpResponse *http.Response

	// yielded values
	httpStatusCode int
	bodyObject     ResponseObject
}

// NewReponse creates a new Response object from the given Session and http.Response.
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

	if !r.bodyObject.HasErrors() {
		if responseHash, ok := r.httpResponse.Header[common.HeaderResponseHash]; ok {
			// Validate the signature
			calculatedHash := signature.HashWithKeys(bodyBytes, []byte(r.session.publicKey), []byte(r.session.privateKey))
			if responseHash[0] != calculatedHash {
				return errors.New(fmt.Sprintf("Signature validation failed. Got %s. Expected %s.", calculatedHash, responseHash[0]))
			}
		} else {
			return errors.New(fmt.Sprintf("%s not present. Unable to validate server response.", common.HeaderResponseHash))
		}
	}

	// all went well
	return nil
}

// HttpResponse gets the http.Response that this Response represents.
func (r *Response) HttpResponse() *http.Response {
	return r.httpResponse
}

// Session gets the Session that was responsible for obtaining this Response.
func (r *Response) Session() *Session {
	return r.session
}

// HttpStatusCode gets the HTTP Status Code of this response.
func (r *Response) HttpStatusCode() int {
	return r.httpResponse.StatusCode
}

// BodyObject gets a real object unmarshalled from the response data.
func (r *Response) BodyObject() ResponseObject {
	return r.bodyObject
}
