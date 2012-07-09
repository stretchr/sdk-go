package stretchr

import (
	"fmt"
)

const (
	DefaultVersion  string = "v1"
	DefaultHost     string = "stretchr.com"
	DefaultProtocol string = "http"
)

// Session represents a Stretchr session.
type Session struct {
	Project    string
	PublicKey  string
	PrivateKey string
	Version    string
	Host       string
	Protocol   string
}

func InProject(project string) *Session {
	s := new(Session)
	s.Project = project
	s.Version = DefaultVersion
	s.Host = DefaultHost
	s.Protocol = DefaultProtocol
	return s
}

func (s *Session) InProject(project string) *Session {
	s.Project = project
	return s
}

func WithKeys(public, private string) *Session {
	s := new(Session)
	s.PublicKey = public
	s.PrivateKey = private
	s.Version = DefaultVersion
	s.Host = DefaultHost
	s.Protocol = DefaultProtocol
	return s
}

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
