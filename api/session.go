package api

import (
	"github.com/stretchr/codecs"
	"github.com/stretchr/codecs/json"
	"github.com/stretchr/sdk-go/common"
	stewstrings "github.com/stretchr/stew/strings"
	"github.com/stretchr/tracer"
)

var Tracer *tracer.Tracer

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
	s.transporter = ActiveLiveTransporter
	s.apiVersion = "1.1"
	s.codec = new(json.JsonCodec)

	Tracer.TraceInfo("New session created. Project: %s. Public Key: %s. Private Key: %s. API Version: %s.", project, publicKey, privateKey, "1")

	return s
}

// BeginTracing starts up the tracer with the specified level
func BeginTracing(level int) {
	Tracer = tracer.New(level)
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
	Tracer.TraceDebug("Setting codec: %s", codec.ContentType())
	s.codec = codec
}

// Transporter gets the current Transporter this Session will use when interacting with
// Stretchr services.
func (s *Session) Transporter() Transporter {
	return s.transporter
}

// SetTransporter sets the Transporter instance to use when interacting with
// Stretchr services.
func (s *Session) SetTransporter(transporter Transporter) *Session {
	s.transporter = transporter
	return s
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

	host := stewstrings.MergeStrings(protocol,
		common.ProtocolSeparator,
		s.project,
		common.HostSeparator,
		common.TopLevelHostName,
		common.ApiVersionPathPrefix,
		s.apiVersion)

	Tracer.TraceInfo("Host string: %s", host)

	return host

}

// At starts a new Request for the specified path.
func (s *Session) At(path string) *Request {
	Tracer.TraceDebug("Creating new Request At: %s", path)

	return NewRequest(s, path)
}
