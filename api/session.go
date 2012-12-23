package api

import (
	"github.com/stretchrcom/codecs"
	"github.com/stretchrcom/codecs/json"
	"github.com/stretchrcom/stretchr-sdk-go/common"
)

// Session provides access to Stretchr services.
type Session struct {
	project     string
	privateKey  string
	publicKey   string
	transporter Transporter
	apiVersion  string
	codec       codecs.Codec
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
	s.codec = new(json.JsonCodec)
	return s
}

// Project gets the project name that this session relates to.
func (s *Session) Project() string {
	return s.project
}

// Codec gets the codec currently being used to communicate with Stretchr.
func (s *Session) Codec() codecs.Codec {
	return s.codec
}

// SetCodec sets the codec to be used to communicate with Stretchr.
func (s *Session) SetCodec(codec codecs.Codec) {
	s.codec = codec
}

// host gets the host to make requests to.
func (s *Session) host() string {

	// get the protocol
	var protocol string
	if s.useSSL {
		protocol = common.HttpProtocolSecure
	} else {
		protocol = common.HttpProtocol
	}

	return MergeStrings(protocol,
		common.ProtocolSeparator,
		s.project,
		common.HostSeparator,
		common.TopLevelHostName,
		common.ApiVersionPathPrefix,
		s.apiVersion)

}

// At starts a new Request for the specified path.
func (s *Session) At(path string) *Request {
	return NewRequest(s, path)
}
