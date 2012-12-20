package stretchr

import (
	"github.com/stretchrcom/testify/assert"
	"github.com/stretchrcom/testify/mock"
	"testing"
)

func TestRequest_Read(t *testing.T) {

	request := NewRequest(getTestSession(), "people")

	returnResponse := new(Response)
	var returnErr error = nil

	mockedTransporter.On("MakeRequest", request).Return(returnResponse, returnErr)
	res, err := request.Read()

	assert.Equal(t, returnErr, err)
	assert.Equal(t, returnResponse, res)

	assert.Equal(t, HttpMethodGet, request.httpMethod)

}

func TestRequest_Delete(t *testing.T) {

	request := NewRequest(getTestSession(), "people")

	returnResponse := new(Response)
	var returnErr error = nil

	mockedTransporter.On("MakeRequest", request).Return(returnResponse, returnErr)
	res, err := request.Delete()

	assert.Equal(t, returnErr, err)
	assert.Equal(t, returnResponse, res)

	assert.Equal(t, HttpMethodDelete, request.httpMethod)

}

func TestRequest_Create(t *testing.T) {

	session := getTestSession()

	returnResponse := new(Response)
	var returnErr error = nil
	mockedTransporter.On("MakeRequest", mock.Anything).Return(returnResponse, returnErr)

	resource := MakeResourceAt("monkey/123")
	resource.Data["name"] = "Mat"
	resource.Data["age"] = 29
	res, err := session.Create(resource)

	assert.Equal(t, res, returnResponse)
	assert.Equal(t, err, returnErr)

	request := mockedTransporter.Calls[0].Arguments.Get(0).(*Request)

	assert.Equal(t, request.httpMethod, HttpMethodPost)
	assert.Equal(t, request.path, resource.ResourcePath())

	expectedBody, _ := ObjectToBytes(resource.Data) //#codecs
	assert.Equal(t, request.body, expectedBody)

}
