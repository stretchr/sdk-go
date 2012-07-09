package stretchr

import (
	"net/http"
)

const (
	ReadMethod    string = "GET"
	CreateMethod  string = "POST"
	UpdateMethod  string = "PUT"
	ReplaceMethod string = "POST"
	DeleteMethod  string = "DELETE"
)

const (
	/*
	   HTTP OK status codes
	*/

	// StatusCodesOKMinimum represents the lowest allowed HTTP status code for requests
	// that are to be considered OK.  Status codes between StatusCodesOKMinimum and StatusCodesOKMaximum
	// will decide the 'Worked' output in the standard response object.
	StatusCodesOKMinimum int = 100

	// StatusCodesOKMaximum represents the highest allowed HTTP status code for requests
	// that are to be considered OK.  Status codes between StatusCodesOKMinimum and StatusCodesOKMaximum
	// will decide the 'Worked' output in the standard response object.
	StatusCodesOKMaximum int = 399
)

// WorkedFromStatusCode gets whether the request was successful based on the given
// HTTP status code.
func WorkedFromStatusCode(statusCode int) bool {
	return statusCode >= StatusCodesOKMinimum && statusCode <= StatusCodesOKMaximum
}

// Requester is an interface describing objects capable of making and processing
// HTTP requests.
type Requester interface {

	// MakeRequest makes a request and returns the response.
	MakeRequest(method, fullUrl, body string) (*StandardResponseObject, *http.Response, error)
}

// DefaultRequester is a Requester object that makes real HTTP requests.
type DefaultRequester struct{}

// MakeRequest makes a request and returns the response.
func (r *DefaultRequester) MakeRequest(method, fullUrl, body string) (*StandardResponseObject, *http.Response, error) {
	panic("Not yet implemented")
	return nil, nil, nil
}

// ActiveRequester is the Requester object this code will use to make requests.
//
// Only change this value if you are writing test code, or if you decide to have more
// control over the actual HTTP requests made by this library.
var ActiveRequester Requester = new(DefaultRequester)
