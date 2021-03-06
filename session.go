package stretchr

import (
	"github.com/stretchr/sdk-go/api"
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
//     Stretchr := stretchr.NewSession(project, apiKey, privateKey)
//
// This enables the following code:
//
//     person, err := Stretchr.ReadOne("people/123")
func NewSession(project, account, apiKey string) *Session {
	s := new(Session)
	s.underlyingSession = api.NewSession(project, account, apiKey)
	return s
}

// At starts a conversation with Stretchr regarding the specified Path.
//
// On its own, At is more or less useless, but as part of a chain of commands,
// is how you specify which resource or set of resources you are referring to.
//
// For example, to delete person with ID `123`, you would use At in the following way:
//
//    Stretchr.At("people/123").Delete()
func (s *Session) At(path string) *Request {
	return NewRequest(s, path)
}

// Project gets the name of the project that this Session interacts with.
func (s *Session) Project() string {
	return s.underlyingSession.Project()
}

// Account gets the name of the project that this Session interacts with.
func (s *Session) Account() string {
	return s.underlyingSession.Account()
}

// SetTransporter sets the Transporter instance to use when interacting with
// Stretchr services.
func (s *Session) SetTransporter(transporter api.Transporter) *Session {
	s.underlyingSession.SetTransporter(transporter)
	return s
}

// SetUseSSL sets whether or not to use SSL
func (s *Session) SetUseSSL(value bool) {
	s.underlyingSession.UseSSL = value
}
