package api

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestLiveTransporter_Interface(t *testing.T) {

	assert.Implements(t, (*Transporter)(nil), new(LiveTransporter))

}

func TestLiveTransporter(t *testing.T) {

}
