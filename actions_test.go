package stretchr

import (
	"github.com/stretchrcom/stretchr-sdk-go/api"
	"github.com/stretchrcom/testify/assert"
	"github.com/stretchrcom/testify/mock"
	"testing"
)

func TestSession_LoadOne(t *testing.T) {

	mockedTransporter := new(api.MockedTransporter)
	api.ActiveLiveTransporter = mockedTransporter

	response := new(api.Response)
	mockedTransporter.On("MakeRequest", mock.Anything).Return(response, nil)

	session := NewSession(TestProjectName, TestPublicKey, TestPrivateKey)

	resource, err := session.LoadOne("people/123")

	if assert.NoError(t, err) {
		assert.NotNil(t, resource)
	}

	mockedTransporter.AssertExpectations(t)

}
