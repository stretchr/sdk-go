package stretchr

import (
	"github.com/stretchr/sdk-go/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSession_NewSession(t *testing.T) {

	s := NewSession("project.company", "123")

	assert.NotNil(t, s, "NewSession")
	assert.Equal(t, "project.company", s.underlyingSession.Project())

}

func TestSession_Project(t *testing.T) {

	s := NewSession("project.company", "123")

	assert.Equal(t, "project.company", s.Project())

}

func TestSession_SetTransporter(t *testing.T) {

	s := NewSession("project.company", "123")

	newTransporter := new(api.LiveTransporter)
	assert.Equal(t, s.SetTransporter(newTransporter), s, "SetTransporter should chain")
	assert.Equal(t, newTransporter, s.underlyingSession.Transporter())

}
