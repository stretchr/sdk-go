package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSession(t *testing.T) {

	s := NewSession("project.company", "123", "456")

	assert.Equal(t, "project.company", s.project)
	assert.Equal(t, "123", s.publicKey)
	assert.Equal(t, "456", s.privateKey)
	assert.Equal(t, ActiveLiveTransporter, s.transporter, "Should default to ActiveLiveTransporter")
	assert.Equal(t, "1.1", s.apiVersion, "apiVersion should default to 1.1")

}

func TestSession_Project(t *testing.T) {

	s := new(Session)
	s.project = "project.company"
	assert.Equal(t, s.project, s.Project())

}

func TestSession_host(t *testing.T) {

	var s *Session

	s = NewSession("project.company", "123", "456")
	s.apiVersion = "2"
	s.useSSL = false
	assert.Equal(t, "http://project.company.stretchr.com/api/v2", s.host())

	s = NewSession("project.company2", "123", "456")
	s.apiVersion = "1"
	s.useSSL = true
	assert.Equal(t, "https://project.company2.stretchr.com/api/v1", s.host())

}

func TestSession_At(t *testing.T) {

	s := new(Session)
	var request *Request = s.At("path")

	if assert.NotNil(t, request) {
		assert.Equal(t, s, request.session)
	}

}

func TestSession_SetTransporter(t *testing.T) {

	s := new(Session)

	newTransporter := new(LiveTransporter)
	assert.Equal(t, s.SetTransporter(newTransporter), s, "SetTransporter should chain")
	assert.Equal(t, newTransporter, s.Transporter())

}
