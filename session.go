package stretchr

// Session provides access to Stretchr services.
type Session struct {
	project     string
	transporter Transporter
}

// NewSession creates a new Session object with the specified project.
func NewSession(project string) *Session {
	s := new(Session)
	s.project = project
	s.transporter = DefaultLiveTransporter
	return s
}

// Project gets the project name that this session relates to.
func (s *Session) Project() string {
	return s.project
}

// On starts a new Request for the specified path.
func (s *Session) On(path string) *Request {
	return NewRequest(s, path)
}
