package stretchr

import (
	"github.com/stretchrcom/stretchr-sdk-go/api"
)

type Session struct {
	session *api.Session
}

func NewSession(project, publicKey, privateKey string) *Session {
	s := new(Session)
	s.session = api.NewSession(project, publicKey, privateKey)
	return s
}

func (s *Session) Project() string {
	return s.session.Project()
}
