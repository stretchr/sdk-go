package stretchr

import (
	"fmt"
)

const (
	DefaultVersion string = "v1"
	DefaultHost string = "stretchr.com"
	DefaultProtocol string = "http"
)

// Session represents a Stretchr session.
type Session struct {
	Project, PublicKey, PrivateKey, Version, Host, Protocol string
}

func InProject(project string) Session {
	return Session{project,"","",DefaultVersion,DefaultHost,DefaultProtocol}
}

func (s Session) InProject(project string) Session {
	s.Project = project
	return s
}

func WithKeys(public, private string) Session {
	return Session{"",public,private,DefaultVersion,DefaultHost,DefaultProtocol}
}

func (s Session) WithKeys(public, private string) Session {
	s.PublicKey = public
	s.PrivateKey = private
	return s
}

func (s Session) baseUrl() string {
	return fmt.Sprintf("%s://%s.%s/api/%s/", s.Protocol, s.Project, s.Host, s.Version)
}

func (s Session) Url(path string) string {
	return fmt.Sprintf("%s%s", s.baseUrl(), path)
}