package stretchr

import (
	"net/http"
)

/*
	Transporter interface
*/

// Transporter describes objects capable of making requests.
type Transporter interface {
	// MakeRequest makes the Request and returns the Response, or an error
	// if there was a problem communicating with the remote server.
	MakeRequest(request *Request) (*Response, error)
}

/*
	LiveTransporter
*/

// DefaultLiveTransporter represents the live transporter instance.
var DefaultLiveTransporter = new(LiveTransporter)

// liveTransporter makes real HTTP requests to remote servers.
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

	return NewResponse(httpResponse), nil
}
