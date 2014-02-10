package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSession(t *testing.T) {

	s := NewSession("project", "company", "123")

	assert.Equal(t, "project", s.project)
	assert.Equal(t, "company", s.account)
	assert.Equal(t, "123", s.apiKey)
	assert.Equal(t, ActiveLiveTransporter, s.transporter, "Should default to ActiveLiveTransporter")
	assert.Equal(t, "1.1", s.apiVersion, "apiVersion should default to 1.1")

}

func TestSession_Project(t *testing.T) {

	s := new(Session)
	s.project = "project.company"
	assert.Equal(t, s.project, s.Project())

}

func TestSession_host(t *testing.T) {

	s := NewSession("project", "company", "123")
	s.apiVersion = "2"
	s.UseSSL = false
	assert.Equal(t, "http://company.stretchr.com/api/v2/project", s.host())

	s = NewSession("project", "company2", "123")
	s.apiVersion = "1"
	s.UseSSL = true
	assert.Equal(t, "https://company2.stretchr.com/api/v1/project", s.host())

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
