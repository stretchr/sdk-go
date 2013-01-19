package api

import (
	"log"
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
var ActiveLiveTransporter Transporter = &LiveTransporter{}

// liveTransporter makes real HTTP requests to remote servers.
type LiveTransporter struct{}

// MakeRequest makes the Request and returns the Response, or an error
// if there was a problem communicating with the remote server.
func (t *LiveTransporter) MakeRequest(request *Request) (*Response, error) {

	// TODO: figure out a way to test this?

	httpRequest, requestErr := request.httpRequest()

	if requestErr != nil {
		return nil, requestErr
	}

	//log.Printf("Making request: %v", httpRequest)

	// make the request
	httpResponse, httpErr := http.DefaultClient.Do(httpRequest)

	if httpErr != nil {
		log.Printf("  Error: %v", httpErr)
		return nil, httpErr
	}

	//log.Printf("  Response: %v", httpResponse)

	response, responseErr := NewResponse(request.session, httpResponse)

	if responseErr != nil {
		return nil, responseErr
	}

	return response, nil
}
