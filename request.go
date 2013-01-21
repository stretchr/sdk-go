package stretchr

import (
	"github.com/stretchrcom/stretchr-sdk-go/api"
)

// Request represents a kind of interaction that can be made against
// the Stretchr services.
type Request struct {
	UnderlyingRequest *api.Request
	session           *Session
}

// NewRequest makes a new Request object with the given session, and path.
//
// It is recommended that you use Session.At to generate a Request instead of calling this
// method directly.
func NewRequest(session *Session, path string) *Request {
	request := new(Request)
	request.UnderlyingRequest = api.NewRequest(session.underlyingSession, path)
	request.session = session
	return request
}

// Session gets the Session object that this request relies on.
func (r *Request) Session() *Session {
	return r.session
}

// Where adds a filter to the request.
func (r *Request) Where(field, match string) *Request {
	r.UnderlyingRequest.Where(field, match)
	return r
}

// Limit sets a limit on the number of resources to get back from Stretchr.
func (r *Request) Limit(value int64) *Request {
	r.UnderlyingRequest.Limit(value)
	return r
}

// Skip sets the number of resources to skip before getting them back from Stretchr.
func (r *Request) Skip(value int64) *Request {
	r.UnderlyingRequest.Skip(value)
	return r
}

func (r *Request) Page(pageNumber, resourcesPerPage int64) *Request {
	r.UnderlyingRequest.Page(pageNumber, resourcesPerPage)
	return r
}

// WithParam sets a query parameter in the request.
func (r *Request) WithParam(key, value string) *Request {
	r.UnderlyingRequest.WithParam(key, value)
	return r
}
