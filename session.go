package stretchr

import (
	"github.com/stretchrcom/stretchr-sdk-go/api"
)

// Session contains project and account information and enables access to
// Stretchr services.
type Session struct {
	// underlyingSession holds the underlying api.Session object.
	underlyingSession *api.Session
}

// NewSession creates a new Session object for interacting with Stretchr services.
//
// Recommended Usage
//
// We recommend that you assign the return of a stretchr.NewSession call to a variable
// called Stretchr for easy reading.
//
// For example,
//
//     Stretchr := stretchr.NewSession(project, publicKey, privateKey)
//
// This enables the following code:
//
//     person, err := Stretchr.ReadOne("people/123")
func NewSession(project, publicKey, privateKey string) *Session {
	s := new(Session)
	s.underlyingSession = api.NewSession(project, publicKey, privateKey)
	return s
}

func (s *Session) At(path string) *Request {
	return NewRequest(s, path)
}

// Project gets the name of the project that this Session interacts with.
func (s *Session) Project() string {
	return s.underlyingSession.Project()
}

// SetTransporter sets the Transporter instance to use when interacting with
// Stretchr services.
func (s *Session) SetTransporter(transporter api.Transporter) *Session {
	s.underlyingSession.SetTransporter(transporter)
	return s
}
