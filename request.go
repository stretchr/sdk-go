package stretchr

import (
	"github.com/stretchrcom/stretchr-sdk-go/api"
)

type Request struct {
	UnderlyingRequest *api.Request
	session           *Session
}

func NewRequest(session *Session, path string) *Request {
	request := new(Request)
	request.UnderlyingRequest = api.NewRequest(session.underlyingSession, path)
	request.session = session
	return request
}

func (r *Request) Session() *Session {
	return r.session
}
