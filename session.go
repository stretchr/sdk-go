package stretchr

import (
	"github.com/stretchrcom/stretchr-sdk-go/api"
)

type Session struct {
	// session holds the underlying api.Session object.
	session *api.Session
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
//     person, err := Stretchr.LoadOne("people/123")
func NewSession(project, publicKey, privateKey string) *Session {
	s := new(Session)
	s.session = api.NewSession(project, publicKey, privateKey)
	return s
}

// Project gets the name of the project that this Session interacts with.
func (s *Session) Project() string {
	return s.session.Project()
}

// SetTransporter sets the Transporter instance to use when interacting with
// Stretchr services.
func (s *Session) SetTransporter(transporter api.Transporter) *Session {
	s.session.SetTransporter(transporter)
	return s
}
