package stretchr

import (
	"github.com/stretchrcom/stretchr-sdk-go/api"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestSession_NewSession(t *testing.T) {

	s := NewSession("project.company", "123", "456")

	assert.NotNil(t, s, "NewSession")
	assert.Equal(t, "project.company", s.underlyingSession.Project())

}

func TestSession_Project(t *testing.T) {

	s := NewSession("project.company", "123", "456")

	assert.Equal(t, "project.company", s.Project())

}

func TestSession_SetTransporter(t *testing.T) {

	s := NewSession("project.company", "123", "456")

	newTransporter := new(api.LiveTransporter)
	assert.Equal(t, s.SetTransporter(newTransporter), s, "SetTransporter should chain")
	assert.Equal(t, newTransporter, s.underlyingSession.Transporter())

}
