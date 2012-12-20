package stretchr

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
	panic("Not implemented")
	return nil, nil
}
