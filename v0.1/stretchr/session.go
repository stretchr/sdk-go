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

func (s *Session) baseUrl() string {
	return fmt.Sprintf("%s://%s.%s/api/%s/", s.Protocol, s.Project, s.Host, s.Version)
}

func (s *Session) Url(path string) string {
	return fmt.Sprintf("%s%s", s.baseUrl(), path)
}
