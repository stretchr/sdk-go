package stretchr

import (
	"fmt"
	"github.com/stretchr/sdk-go/api"
	"github.com/stretchr/sdk-go/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

/*
	ReadOne
*/

func TestRequest_ReadOne(t *testing.T) {

	mockedTransporter := new(api.MockedTransporter)
	api.ActiveLiveTransporter = mockedTransporter

	// make a response
	response := NewTestResponse(200, map[string]interface{}{"name": "Mat"}, nil, "", nil)
	mockedTransporter.On("MakeRequest", mock.Anything).Return(response, nil)

	session := NewSession(TestProjectName, TestPublicKey, TestPrivateKey)

	resource, err := session.At("people/123").ReadOne()

	if assert.NoError(t, err) {
		assert.NotNil(t, resource)
	}

	mockedTransporter.AssertExpectations(t)
	request := mockedTransporter.Calls[0].Arguments[0].(*api.Request)

	assert.Equal(t, request.HttpMethod(), common.HttpMethodGet)
	assert.Equal(t, request.Path(), "people/123")
	assert.Equal(t, request.Body(), []byte(""))

	assert.Equal(t, resource.ResourceData()["name"], response.BodyObject().Data().(map[string]interface{})["name"])
	assert.Equal(t, resource.ResourcePath(), "people/123")

}

func TestRequest_ReadOne_ReadError(t *testing.T) {

	mockedTransporter := new(api.MockedTransporter)
	api.ActiveLiveTransporter = mockedTransporter

	// make a response
	mockedTransporter.On("MakeRequest", mock.Anything).Return(nil, assert.AnError)

	session := NewSession(TestProjectName, TestPublicKey, TestPrivateKey)

	resource, err := session.At("people/123").ReadOne()

	if assert.Nil(t, resource) {
		assert.Equal(t, assert.AnError, err)
	}

}

func TestRequest_ReadOne_StretchrError(t *testing.T) {

	mockedTransporter := new(api.MockedTransporter)
	api.ActiveLiveTransporter = mockedTransporter

	// make a response
	response := NewTestResponse(500, nil, []map[string]interface{}{map[string]interface{}{"~m": "Something went wrong"}}, "", nil)
	mockedTransporter.On("MakeRequest", mock.Anything).Return(response, nil)

	session := NewSession(TestProjectName, TestPublicKey, TestPrivateKey)

	resource, err := session.At("people/123").ReadOne()

	if assert.Nil(t, resource) {
		assert.Equal(t, "Something went wrong", fmt.Sprintf("%s", err))
	}

}

/*
	ReadMany
*/

func TestRequest_ReadMany(t *testing.T) {

	mockedTransporter := new(api.MockedTransporter)
	api.ActiveLiveTransporter = mockedTransporter

	// make a response
	response := NewTestResponse(200, []map[string]interface{}{map[string]interface{}{"name": "Mat", common.DataFieldID: "ABC"},
		map[string]interface{}{"name": "Tyler", common.DataFieldID: "DEF"}}, nil, "", nil)
	mockedTransporter.On("MakeRequest", mock.Anything).Return(response, nil)

	session := NewSession(TestProjectName, TestPublicKey, TestPrivateKey)

	resourceCollection, err := session.At("people").ReadMany()

	if assert.NoError(t, err) {
		assert.Equal(t, 2, len(resourceCollection.Resources))
	}

	mockedTransporter.AssertExpectations(t)
	request := mockedTransporter.Calls[0].Arguments[0].(*api.Request)

	assert.Equal(t, request.HttpMethod(), common.HttpMethodGet)
	assert.Equal(t, request.Path(), "people")
	assert.Equal(t, request.Body(), []byte(""))

	resource1 := resourceCollection.Resources[0]
	resource2 := resourceCollection.Resources[1]

	assert.Equal(t, resource1.ResourceData()["name"], response.BodyObject().Data().([]interface{})[0].(map[string]interface{})["name"])
	assert.Equal(t, resource2.ResourceData()["name"], response.BodyObject().Data().([]interface{})[1].(map[string]interface{})["name"])
	assert.Equal(t, resource1.ResourcePath(), "people/ABC")
	assert.Equal(t, resource2.ResourcePath(), "people/DEF")
	assert.Equal(t, resource1.ResourcePath(), "people/ABC")
	assert.Equal(t, resource2.ResourcePath(), "people/DEF")

}

func TestRequest_ReadMany_ReadError(t *testing.T) {

	mockedTransporter := new(api.MockedTransporter)
	api.ActiveLiveTransporter = mockedTransporter

	// make a response
	mockedTransporter.On("MakeRequest", mock.Anything).Return(nil, assert.AnError)

	session := NewSession(TestProjectName, TestPublicKey, TestPrivateKey)

	resource, err := session.At("people/123").ReadMany()

	if assert.Nil(t, resource) {
		assert.Equal(t, assert.AnError, err)
	}

}

func TestRequest_ReadMany_StretchrError(t *testing.T) {

	mockedTransporter := new(api.MockedTransporter)
	api.ActiveLiveTransporter = mockedTransporter

	// make a response
	response := NewTestResponse(500, nil, []map[string]interface{}{map[string]interface{}{"~m": "Something went wrong"}}, "", nil)
	mockedTransporter.On("MakeRequest", mock.Anything).Return(response, nil)

	session := NewSession(TestProjectName, TestPublicKey, TestPrivateKey)

	resource, err := session.At("people/123").ReadMany()

	if assert.Nil(t, resource) {
		assert.Equal(t, "Something went wrong", fmt.Sprintf("%s", err))
	}

}

/*
	Delete
*/
func TestRequest_Delete(t *testing.T) {

	mockedTransporter := new(api.MockedTransporter)
	api.ActiveLiveTransporter = mockedTransporter

	// make a response
	response := NewTestResponse(200, nil, nil, "", api.ChangeInfo(map[string]interface{}{api.ChangeInfoPublicFieldDeleted: 5}))
	mockedTransporter.On("MakeRequest", mock.Anything).Return(response, nil)

	session := NewSession(TestProjectName, TestPublicKey, TestPrivateKey)

	changeInfo, err := session.At("people/123").Delete()

	if assert.NoError(t, err) {
		assert.NotNil(t, changeInfo)
	}

	mockedTransporter.AssertExpectations(t)
	request := mockedTransporter.Calls[0].Arguments[0].(*api.Request)

	assert.Equal(t, request.HttpMethod(), common.HttpMethodDelete)
	assert.Equal(t, request.Path(), "people/123")
	assert.Equal(t, request.Body(), []byte(""))

	assert.Equal(t, changeInfo.Deleted(), 5)

}

func TestRequest_Delete_ReadError(t *testing.T) {

	mockedTransporter := new(api.MockedTransporter)
	api.ActiveLiveTransporter = mockedTransporter

	// make a response
	mockedTransporter.On("MakeRequest", mock.Anything).Return(nil, assert.AnError)

	session := NewSession(TestProjectName, TestPublicKey, TestPrivateKey)

	resource, err := session.At("people/123").Delete()

	if assert.Nil(t, resource) {
		assert.Equal(t, assert.AnError, err)
	}

}

func TestRequest_Delete_StretchrError(t *testing.T) {

	mockedTransporter := new(api.MockedTransporter)
	api.ActiveLiveTransporter = mockedTransporter

	// make a response
	response := NewTestResponse(500, nil, []map[string]interface{}{map[string]interface{}{"~m": "Something went wrong"}}, "", nil)
	mockedTransporter.On("MakeRequest", mock.Anything).Return(response, nil)

	session := NewSession(TestProjectName, TestPublicKey, TestPrivateKey)

	resource, err := session.At("people/123").Delete()

	if assert.Nil(t, resource) {
		assert.Equal(t, "Something went wrong", fmt.Sprintf("%s", err))
	}

}

func TestRequest_Create(t *testing.T) {

	resource := MakeResourceAt("people")
	resource.Set("name", "Mat").Set("age", 29)

	mockedTransporter := new(api.MockedTransporter)
	api.ActiveLiveTransporter = mockedTransporter

	// make a response
	response := NewTestResponse(200, nil, nil, "", api.ChangeInfo(map[string]interface{}{api.ChangeInfoPublicFieldCreated: 1, api.ChangeInfoPublicFieldDeltas: []interface{}{map[string]interface{}{common.DataFieldID: "hello", api.ChangeInfoPublicFieldDeltaCreated: 123}}}))
	mockedTransporter.On("MakeRequest", mock.Anything).Return(response, nil)

	session := NewSession(TestProjectName, TestPublicKey, TestPrivateKey)

	changeInfo, err := session.At("people").Create(resource)

	if assert.NoError(t, err) {
		if assert.NotNil(t, changeInfo) {
			mockedTransporter.AssertExpectations(t)
			request := mockedTransporter.Calls[0].Arguments[0].(*api.Request)

			assert.Equal(t, request.HttpMethod(), common.HttpMethodPost)
			assert.Equal(t, request.Path(), "people")
			assert.Equal(t, changeInfo.Created(), 1)
			assert.Equal(t, resource.ID(), "hello")
			assert.Equal(t, resource.ResourceData()[api.ChangeInfoPublicFieldDeltaCreated].(float64), 123)
		}
	}

}

func TestRequest_Save_Create(t *testing.T) {

	resource := MakeResourceAt("people")
	resource.Set("name", "Mat").Set("age", 29)

	mockedTransporter := new(api.MockedTransporter)
	api.ActiveLiveTransporter = mockedTransporter

	// make a response
	response := NewTestResponse(200, nil, nil, "", api.ChangeInfo(map[string]interface{}{api.ChangeInfoPublicFieldCreated: 1, api.ChangeInfoPublicFieldDeltas: []interface{}{map[string]interface{}{common.DataFieldID: "hello", api.ChangeInfoPublicFieldDeltaCreated: 123}}}))
	mockedTransporter.On("MakeRequest", mock.Anything).Return(response, nil)

	session := NewSession(TestProjectName, TestPublicKey, TestPrivateKey)

	changeInfo, err := session.At("people").Save(resource)

	if assert.NoError(t, err) {
		if assert.NotNil(t, changeInfo) {
			mockedTransporter.AssertExpectations(t)
			request := mockedTransporter.Calls[0].Arguments[0].(*api.Request)

			assert.Equal(t, request.HttpMethod(), common.HttpMethodPut)
			assert.Equal(t, request.Path(), "people")
			assert.Equal(t, changeInfo.Created(), 1)
		}
	}

}

func TestRequest_Update_Create(t *testing.T) {

	resource := MakeResourceAt("people")
	resource.Set("name", "Mat").Set("age", 29)

	mockedTransporter := new(api.MockedTransporter)
	api.ActiveLiveTransporter = mockedTransporter

	// make a response
	response := NewTestResponse(200, nil, nil, "", api.ChangeInfo(map[string]interface{}{api.ChangeInfoPublicFieldUpdated: 1, api.ChangeInfoPublicFieldDeltas: []interface{}{map[string]interface{}{common.DataFieldID: "hello", api.ChangeInfoPublicFieldDeltaCreated: 123}}}))
	mockedTransporter.On("MakeRequest", mock.Anything).Return(response, nil)

	session := NewSession(TestProjectName, TestPublicKey, TestPrivateKey)

	changeInfo, err := session.At("people").Update(resource)

	if assert.NoError(t, err) {
		if assert.NotNil(t, changeInfo) {
			mockedTransporter.AssertExpectations(t)
			request := mockedTransporter.Calls[0].Arguments[0].(*api.Request)

			assert.Equal(t, request.HttpMethod(), common.HttpMethodPut)
			assert.Equal(t, request.Path(), "people")
			assert.Equal(t, changeInfo.Updated(), 1)
		}
	}

}

func TestRequest_Replace_Create(t *testing.T) {

	resource := MakeResourceAt("people")
	resource.Set("name", "Mat").Set("age", 29)

	mockedTransporter := new(api.MockedTransporter)
	api.ActiveLiveTransporter = mockedTransporter

	// make a response
	response := NewTestResponse(200, nil, nil, "", api.ChangeInfo(map[string]interface{}{api.ChangeInfoPublicFieldCreated: 1, api.ChangeInfoPublicFieldDeltas: []interface{}{map[string]interface{}{common.DataFieldID: "new"}}}))
	mockedTransporter.On("MakeRequest", mock.Anything).Return(response, nil)

	session := NewSession(TestProjectName, TestPublicKey, TestPrivateKey)

	changeInfo, err := session.At("people").Replace(resource)

	if assert.NoError(t, err) {
		if assert.NotNil(t, changeInfo) {
			mockedTransporter.AssertExpectations(t)
			request := mockedTransporter.Calls[0].Arguments[0].(*api.Request)

			assert.Equal(t, request.HttpMethod(), common.HttpMethodPost)
			assert.Equal(t, request.Path(), "people")
			assert.Equal(t, changeInfo.Created(), 1)
		}
	}

}

func TestRequest_Replace_Replace(t *testing.T) {

	resource := MakeResourceAt("people")
	resource.Set("name", "Mat").Set("age", 29)

	mockedTransporter := new(api.MockedTransporter)
	api.ActiveLiveTransporter = mockedTransporter

	// make a response
	response := NewTestResponse(200, nil, nil, "", api.ChangeInfo(map[string]interface{}{api.ChangeInfoPublicFieldCreated: 1, api.ChangeInfoPublicFieldDeltas: []interface{}{map[string]interface{}{common.DataFieldID: "hello", api.ChangeInfoPublicFieldDeltaCreated: 123}}}))
	mockedTransporter.On("MakeRequest", mock.Anything).Return(response, nil)

	session := NewSession(TestProjectName, TestPublicKey, TestPrivateKey)

	changeInfo, err := session.At("people").Create(resource)

	if assert.NoError(t, err) {
		if assert.NotNil(t, changeInfo) {
			mockedTransporter.AssertExpectations(t)
			request := mockedTransporter.Calls[0].Arguments[0].(*api.Request)

			assert.Equal(t, request.HttpMethod(), common.HttpMethodPost)
			assert.Equal(t, request.Path(), "people")
			assert.Equal(t, changeInfo.Created(), 1)
		}
	}

}

/*
	Save
*/

func TestRequest_Update_Update(t *testing.T) {

	resource := MakeResourceAt("people")
	resource.Set("name", "Mat").Set("age", 29)

	mockedTransporter := new(api.MockedTransporter)
	api.ActiveLiveTransporter = mockedTransporter

	// make a response
	response := NewTestResponse(200, nil, nil, "", api.ChangeInfo(map[string]interface{}{api.ChangeInfoPublicFieldCreated: 1, api.ChangeInfoPublicFieldDeltas: []interface{}{map[string]interface{}{common.DataFieldID: "hello", api.ChangeInfoPublicFieldDeltaCreated: 123}}}))
	mockedTransporter.On("MakeRequest", mock.Anything).Return(response, nil)

	session := NewSession(TestProjectName, TestPublicKey, TestPrivateKey)

	changeInfo, err := session.At("people").Update(resource)

	if assert.NoError(t, err) {
		if assert.NotNil(t, changeInfo) {
			mockedTransporter.AssertExpectations(t)
			request := mockedTransporter.Calls[0].Arguments[0].(*api.Request)

			assert.Equal(t, request.HttpMethod(), common.HttpMethodPut)
			assert.Equal(t, request.Path(), "people")
			assert.Equal(t, changeInfo.Created(), 1)
		}
	}

}

func TestRequest_Save_Update(t *testing.T) {

	resource := MakeResourceAt("people")
	resource.Set("name", "Mat").Set("age", 29)

	mockedTransporter := new(api.MockedTransporter)
	api.ActiveLiveTransporter = mockedTransporter

	// make a response
	response := NewTestResponse(200, nil, nil, "", api.ChangeInfo(map[string]interface{}{api.ChangeInfoPublicFieldUpdated: 1, api.ChangeInfoPublicFieldDeltas: []interface{}{map[string]interface{}{common.DataFieldID: "hello", api.ChangeInfoPublicFieldDeltaCreated: 123}}}))
	mockedTransporter.On("MakeRequest", mock.Anything).Return(response, nil)

	session := NewSession(TestProjectName, TestPublicKey, TestPrivateKey)

	changeInfo, err := session.At("people").Save(resource)

	if assert.NoError(t, err) {
		if assert.NotNil(t, changeInfo) {
			mockedTransporter.AssertExpectations(t)
			request := mockedTransporter.Calls[0].Arguments[0].(*api.Request)

			assert.Equal(t, request.HttpMethod(), common.HttpMethodPut)
			assert.Equal(t, request.Path(), "people")
			assert.Equal(t, changeInfo.Updated(), 1)
		}
	}

}
