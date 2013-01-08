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

	// make a response
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

func TestSession_LoadMany(t *testing.T) {

	mockedTransporter := new(api.MockedTransporter)
	api.ActiveLiveTransporter = mockedTransporter

	// make a response
	response := NewTestResponse(200, []map[string]interface{}{map[string]interface{}{"name": "Mat"},
		map[string]interface{}{"name": "Tyler"}}, nil, "", nil)
	mockedTransporter.On("MakeRequest", mock.Anything).Return(response, nil)

	session := NewSession(TestProjectName, TestPublicKey, TestPrivateKey)

	resources, err := session.LoadMany("people")

	if assert.NoError(t, err) {
		assert.Equal(t, 2, len(resources))
	}

	mockedTransporter.AssertExpectations(t)
	request := mockedTransporter.Calls[0].Arguments[0].(*api.Request)

	assert.Equal(t, request.HttpMethod(), common.HttpMethodGet)
	assert.Equal(t, request.Path(), "people")
	assert.Equal(t, request.Body(), []byte(""))

	resource1 := resources[0]
	resource2 := resources[1]

	assert.Equal(t, resource1.ResourceData()["name"], response.BodyObject().Data().([]interface{})[0].(map[string]interface{})["name"])
	assert.Equal(t, resource2.ResourceData()["name"], response.BodyObject().Data().([]interface{})[1].(map[string]interface{})["name"])

}
