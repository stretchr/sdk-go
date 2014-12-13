package api

import (
	"net/http"
)

/*
	Transporter interface
*/

// Transporter describes objects capable of making HTTP requests to Stretchr.
//
// The LiveTransporter object makes real HTTP requests, where the MockedTransporter
// allows you to write tests against your Stretchr code.
type Transporter interface {
	// MakeRequest makes the Request and returns the Response, or an error
	// if there was a problem communicating with the remote server.
	MakeRequest(request *Request) (*Response, error)
}

/*
	LiveTransporter
*/

// ActiveLiveTransporter represents the live transporter instance.
var ActiveLiveTransporter = &LiveTransporter{}

// LiveTransporter makes real HTTP requests to remote servers.
type LiveTransporter struct{}

// MakeRequest makes the Request and returns the Response, or an error
// if there was a problem communicating with the remote server.
func (t *LiveTransporter) MakeRequest(request *Request) (*Response, error) {

	httpRequest, requestErr := request.httpRequest()

	if requestErr != nil {
		return nil, requestErr
	}

	// make the request
	httpResponse, httpErr := http.DefaultClient.Do(httpRequest)

	if httpErr != nil {
		return nil, httpErr
	}

	response, responseErr := NewResponse(request.session, httpResponse)

	if responseErr != nil {
		return nil, responseErr
	}

	return response, nil
}
