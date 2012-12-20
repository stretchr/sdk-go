package stretchr

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestNewSession(t *testing.T) {

	s := NewSession("project.company")

	assert.Equal(t, "project.company", s.project)
	assert.Equal(t, DefaultLiveTransporter, s.transporter, "Should default to DefaultLiveTransporter")

}

func TestSession_Project(t *testing.T) {

	s := new(Session)
	s.project = "project.company"
	assert.Equal(t, s.project, s.Project())

}

func TestSession_On(t *testing.T) {

	s := new(Session)
	var request *Request = s.On("path")

	if assert.NotNil(t, request) {
		assert.Equal(t, s, request.session)
	}

}
