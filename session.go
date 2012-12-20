package stretchr

// Session provides access to Stretchr services.
type Session struct {
	project     string
	privateKey  string
	publicKey   string
	transporter Transporter
	apiVersion  string
	useSSL      bool
}

// NewSession creates a new Session object with the specified project.
func NewSession(project, publicKey, privateKey string) *Session {
	s := new(Session)
	s.project = project
	s.publicKey = publicKey
	s.privateKey = privateKey
	s.transporter = DefaultLiveTransporter
	s.apiVersion = "1"
	return s
}

// Project gets the project name that this session relates to.
func (s *Session) Project() string {
	return s.project
}

// host gets the host to make requests to.
func (s *Session) host() string {

	// get the protocol
	var protocol string
	if s.useSSL {
		protocol = httpProtocolSecure
	} else {
		protocol = httpProtocol
	}

	return MergeStrings(protocol,
		protocolSeparator,
		s.project,
		hostSeparator,
		topLevelHostName,
		apiVersionPathPrefix,
		s.apiVersion)

}

// At starts a new Request for the specified path.
func (s *Session) At(path string) *Request {
	return NewRequest(s, path)
}
