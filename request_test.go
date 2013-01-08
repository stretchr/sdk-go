package stretchr

import (
	"github.com/stretchrcom/stretchr-sdk-go/api"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

var testSession *Session
var mockedTransporter *api.MockedTransporter

func getTestSession() *Session {
	testSession = NewSession("test", "123", "456")
	testSession.underlyingSession = api.NewSession("project", "publicKey", "privateKey")
	mockedTransporter = new(api.MockedTransporter)
	testSession.SetTransporter(mockedTransporter)
	return testSession
}

func TestNewRequest(t *testing.T) {

	r := NewRequest(getTestSession(), "people/123")

	if assert.NotNil(t, r) {
		assert.Equal(t, r.session, testSession)
		assert.Equal(t, r.UnderlyingRequest.Session(), testSession.underlyingSession)
		assert.Equal(t, r.UnderlyingRequest.Path(), "people/123")
	}

}

func TestSession(t *testing.T) {

	r := NewRequest(getTestSession(), "people/123")
	assert.Equal(t, r.Session(), testSession)

}
