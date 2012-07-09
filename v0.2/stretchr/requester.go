package stretchr

import (
	"net/http"
)

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
