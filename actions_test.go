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

func TestRequest_CreateMany(t *testing.T) {

	request := NewRequest(getTestSession(), "monkey")

	returnResponse := new(Response)
	var returnErr error = nil
	mockedTransporter.On("MakeRequest", request).Return(returnResponse, returnErr)

	resource1 := MakeResourceAt("people/123")
	resource1.Data["name"] = "Mat"
	resource1.Data["age"] = 29

	resource2 := MakeResourceAt("people/124")
	resource2.Data["name"] = "Tyler"
	resource2.Data["age"] = 28

	resourceCollection := []Resource{resource1, resource2}

	res, err := request.CreateMany(resourceCollection)

	assert.Equal(t, res, returnResponse)
	assert.Equal(t, err, returnErr)

	request = mockedTransporter.Calls[0].Arguments.Get(0).(*Request)

	assert.Equal(t, request.httpMethod, HttpMethodPost)
	assert.Equal(t, request.path, "monkey")

	expectedBody, _ := ObjectToBytes([]interface{}{resource1.Data, resource2.Data}) //#codecs
	assert.Equal(t, request.body, expectedBody)

}

func TestRequest_Update(t *testing.T) {

	session := getTestSession()

	returnResponse := new(Response)
	var returnErr error = nil
	mockedTransporter.On("MakeRequest", mock.Anything).Return(returnResponse, returnErr)

	resource := MakeResourceAt("monkey/123")
	resource.Data["name"] = "Mat"
	resource.Data["age"] = 29
	res, err := session.Update(resource)

	assert.Equal(t, res, returnResponse)
	assert.Equal(t, err, returnErr)

	request := mockedTransporter.Calls[0].Arguments.Get(0).(*Request)

	assert.Equal(t, request.httpMethod, HttpMethodPut)
	assert.Equal(t, request.path, resource.ResourcePath())

	expectedBody, _ := ObjectToBytes(resource.Data) //#codecs
	assert.Equal(t, request.body, expectedBody)

}

func TestRequest_Replace(t *testing.T) {

	session := getTestSession()

	returnResponse := new(Response)
	var returnErr error = nil
	mockedTransporter.On("MakeRequest", mock.Anything).Return(returnResponse, returnErr)

	resource := MakeResourceAt("monkey/123")
	resource.Data["name"] = "Mat"
	resource.Data["age"] = 29
	res, err := session.Replace(resource)

	assert.Equal(t, res, returnResponse)
	assert.Equal(t, err, returnErr)

	request := mockedTransporter.Calls[0].Arguments.Get(0).(*Request)

	assert.Equal(t, request.httpMethod, HttpMethodPost)
	assert.Equal(t, request.path, resource.ResourcePath())

	expectedBody, _ := ObjectToBytes(resource.Data) //#codecs
	assert.Equal(t, request.body, expectedBody)

}
