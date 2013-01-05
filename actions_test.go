package stretchr

import (
	"github.com/stretchrcom/stretchr-sdk-go/api"
	"github.com/stretchrcom/stretchr-sdk-go/common"
	"github.com/stretchrcom/testify/assert"
	"github.com/stretchrcom/testify/mock"
	"testing"
)

func TestSession_LoadOne(t *testing.T) {

	mockedTransporter := new(api.MockedTransporter)
	api.ActiveLiveTransporter = mockedTransporter

	response := NewTestResponse(200, map[string]interface{}{"name": "Mat"}, nil, "", nil)
	mockedTransporter.On("MakeRequest", mock.Anything).Return(response, nil)

	session := NewSession(TestProjectName, TestPublicKey, TestPrivateKey)

	resource, err := session.LoadOne("people/123")

	if assert.NoError(t, err) {
		assert.NotNil(t, resource)
	}

	mockedTransporter.AssertExpectations(t)
	request := mockedTransporter.Calls[0].Arguments[0].(*api.Request)

	assert.Equal(t, request.HttpMethod(), common.HttpMethodGet)
	assert.Equal(t, request.Path(), "people/123")
	assert.Equal(t, request.Body(), []byte(""))

	assert.Equal(t, resource.ResourceData()["name"], response.BodyObject().Data().(map[string]interface{})["name"])

}
