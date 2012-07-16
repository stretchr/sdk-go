package stretchr

import (
	"fmt"
)

const (
	// DefaultVersion is the default version of the Stretchr service.
	DefaultVersion string = "v1"

	// DefaultHost is the default host of the Stretchr service.
	DefaultHost string = "stretchr.com"

	// DefaultProtocol is the default protocol of the Stretchr service.
	DefaultProtocol string = "http"
)

// Session represents a Stretchr session.
type Session struct {

	// Project is the name of the project this Session represents.
	Project string

	// PublicKey is the public key of the account that this Session represents.
	PublicKey string

	// PrivateKey is the private key of the account that this Session represents.
	PrivateKey string

	// Version is the string representing the version of the Stretchr service to use (see DefaultVersion).
	Version string

	// Host is the string representing the host of the Stretchr service to use (see DefaultHost).
	Host string

	// Protocol is the string representing the protocol by which to access the Stretchr service to use (see DefaultProtocol).
	Protocol string
}

/*
	Construction and configuration
*/

// InProject creates a new Session and sets the Project field.
func InProject(project string) *Session {
	s := new(Session)
	s.Project = project
	s.Version = DefaultVersion
	s.Host = DefaultHost
	s.Protocol = DefaultProtocol
	return s
}

// InProject sets the Project field of this Session.
func (s *Session) InProject(project string) *Session {
	s.Project = project
	return s
}

// WithKeys creates a new session with the specified keys.
func WithKeys(public, private string) *Session {
	s := new(Session)
	s.PublicKey = public
	s.PrivateKey = private
	s.Version = DefaultVersion
	s.Host = DefaultHost
	s.Protocol = DefaultProtocol
	return s
}

// WithKeys sets the keys for this Session.
func (s *Session) WithKeys(public, private string) *Session {
	s.PublicKey = public
	s.PrivateKey = private
	return s
}

/*
	Resources
*/

// MakeResource makes a resource with the given path.
//
// 	resource := session.MakeResource("people")
func (s *Session) MakeResource(path string) *Resource {
	return MakeResource(s, path)
}

// MakeResourcef makes a resource with the given path allowing for formatting.
//
// 	resource := session.MakeResource("people/%s/books", personID)
func (s *Session) MakeResourcef(path string, args ...interface{}) *Resource {
	return s.MakeResource(fmt.Sprintf(path, args...))
}

/*
	URLs
*/

// baseUrl gets the base URL for requests based on the settings in this session.
func (s *Session) baseUrl() string {
	return fmt.Sprintf("%s://%s.%s/api/%s/", s.Protocol, s.Project, s.Host, s.Version)
}

// Url gets the absolute URL for the specified path based on the settings in 
// this session.
func (s *Session) Url(path string) string {
	return fmt.Sprintf("%s%s", s.baseUrl(), path)
}

/*
	Many
*/

// Many starts a Many object capable of reading or deleting multiple resources.
func (s *Session) Many(path string) *Many {
	return makeMany(s, path)
}

// Manyf starts a Many object capable of reading or deleting multiple resources
// but also allows formatting in the path.
//
// For example:
//  s.Manyf("people/%s/books", personID)
func (s *Session) Manyf(path string, args ...interface{}) *Many {
	return s.Many(fmt.Sprintf(path, args...))
}
