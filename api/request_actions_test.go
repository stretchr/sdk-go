package api

import (
	"testing"

	"github.com/stretchr/sdk-go/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRequest_Read(t *testing.T) {

	request := NewRequest(getTestSession(), "people")

	returnResponse := new(Response)
	var returnErr error

	mockedTransporter.On("MakeRequest", request).Return(returnResponse, returnErr)
	res, err := request.Read()

	assert.Equal(t, returnErr, err)
	assert.Equal(t, returnResponse, res)

	assert.Equal(t, common.HTTPMethodGet, request.httpMethod)

}

func TestRequest_Delete(t *testing.T) {

	request := NewRequest(getTestSession(), "people")

	returnResponse := new(Response)
	var returnErr error

	mockedTransporter.On("MakeRequest", request).Return(returnResponse, returnErr)
	res, err := request.Delete()

	assert.Equal(t, returnErr, err)
	assert.Equal(t, returnResponse, res)

	assert.Equal(t, common.HTTPMethodDelete, request.httpMethod)

}

func TestRequest_Create(t *testing.T) {

	request := NewRequest(getTestSession(), "monkey/123")

	returnResponse := new(Response)
	var returnErr error
	mockedTransporter.On("MakeRequest", mock.Anything).Return(returnResponse, returnErr)

	resource := MakeTestResourceAt("monkey/123")
	resource.Data["name"] = "Mat"
	resource.Data["age"] = 29
	res, err := request.Create(resource)

	assert.Equal(t, res, returnResponse)
	assert.Equal(t, err, returnErr)

	request = mockedTransporter.Calls[0].Arguments.Get(0).(*Request)

	assert.Equal(t, request.httpMethod, common.HTTPMethodPost)
	assert.Equal(t, request.path, resource.ResourcePath())

	expectedBody, _ := request.session.codec.Marshal(resource.Data, nil)
	assert.Equal(t, request.body, expectedBody)

}

func TestRequest_CreateMany(t *testing.T) {

	request := NewRequest(getTestSession(), "monkey")

	returnResponse := new(Response)
	var returnErr error
	mockedTransporter.On("MakeRequest", request).Return(returnResponse, returnErr)

	resource1 := MakeTestResourceAt("people/123")
	resource1.Data["name"] = "Mat"
	resource1.Data["age"] = 29

	resource2 := MakeTestResourceAt("people/124")
	resource2.Data["name"] = "Tyler"
	resource2.Data["age"] = 28

	resourceCollection := []Resource{resource1, resource2}

	res, err := request.CreateMany(resourceCollection)

	assert.Equal(t, res, returnResponse)
	assert.Equal(t, err, returnErr)

	request = mockedTransporter.Calls[0].Arguments.Get(0).(*Request)

	assert.Equal(t, request.httpMethod, common.HTTPMethodPost)
	assert.Equal(t, request.path, "monkey")

	expectedBody, _ := request.session.codec.Marshal([]interface{}{resource1.Data, resource2.Data}, nil)
	assert.Equal(t, request.body, expectedBody)

}

func TestRequest_Update(t *testing.T) {

	request := NewRequest(getTestSession(), "monkey/123")

	returnResponse := new(Response)
	var returnErr error
	mockedTransporter.On("MakeRequest", mock.Anything).Return(returnResponse, returnErr)

	resource := MakeTestResourceAt("monkey/123")
	resource.Data["name"] = "Mat"
	resource.Data["age"] = 29
	res, err := request.Update(resource)

	assert.Equal(t, res, returnResponse)
	assert.Equal(t, err, returnErr)

	request = mockedTransporter.Calls[0].Arguments.Get(0).(*Request)

	assert.Equal(t, request.httpMethod, common.HTTPMethodPatch)
	assert.Equal(t, request.path, resource.ResourcePath())

	expectedBody, _ := request.session.codec.Marshal(resource.Data, nil)
	assert.Equal(t, request.body, expectedBody)

}

func TestRequest_Replace(t *testing.T) {

	request := NewRequest(getTestSession(), "monkey/123")

	returnResponse := new(Response)
	var returnErr error
	mockedTransporter.On("MakeRequest", mock.Anything).Return(returnResponse, returnErr)

	resource := MakeTestResourceAt("monkey/123")
	resource.Data["name"] = "Mat"
	resource.Data["age"] = 29
	res, err := request.Replace(resource)

	assert.Equal(t, res, returnResponse)
	assert.Equal(t, err, returnErr)

	request = mockedTransporter.Calls[0].Arguments.Get(0).(*Request)

	assert.Equal(t, request.httpMethod, common.HTTPMethodPut)
	assert.Equal(t, request.path, resource.ResourcePath())

	expectedBody, _ := request.session.codec.Marshal(resource.Data, nil)
	assert.Equal(t, request.body, expectedBody)

}
