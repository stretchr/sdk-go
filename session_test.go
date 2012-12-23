package stretchr

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestSession_NewSession(t *testing.T) {

	s := NewSession("project.company", "123", "456")

	assert.NotNil(t, s, "NewSession")
	assert.Equal(t, "project.company", s.session.Project())

}

func TestSession_Project(t *testing.T) {

	s := NewSession("project.company", "123", "456")

	assert.Equal(t, "project.company", s.Project())

}
