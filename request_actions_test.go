package stretchr

import (
	"fmt"
	"testing"

	"github.com/stretchr/sdk-go/api"
	"github.com/stretchr/sdk-go/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/*
	ReadOne
*/

func TestRequest_ReadOne(t *testing.T) {

	mockedTransporter := new(api.MockedTransporter)
	api.ActiveLiveTransporter = mockedTransporter

	// make a response
	response := NewTestResponse(200, map[string]interface{}{"~data": []interface{}{map[string]interface{}{"name": "Mat"}}}, nil, "", nil)
	mockedTransporter.On("MakeRequest", mock.Anything).Return(response, nil)

	session := NewSession(TestProjectName, TestAccountName, TestAPIKey)

	resource, err := session.At("people/123").ReadOne()

	if assert.NoError(t, err) {
		assert.NotNil(t, resource)
	}

	mockedTransporter.AssertExpectations(t)
	request := mockedTransporter.Calls[0].Arguments[0].(*api.Request)

	assert.Equal(t, request.HTTPMethod(), common.HTTPMethodGet)
	assert.Equal(t, request.Path(), "people/123")
	assert.Equal(t, request.Body(), []byte(""))

	assert.Equal(t, resource.ResourceData()["name"], response.BodyObject().Data().(map[string]interface{})["~data"].([]interface{})[0].(map[string]interface{})["name"])
	assert.Equal(t, resource.ResourcePath(), "people/123")

}

func TestRequest_ReadOne_ReadError(t *testing.T) {

	mockedTransporter := new(api.MockedTransporter)
	api.ActiveLiveTransporter = mockedTransporter

	// make a response
	mockedTransporter.On("MakeRequest", mock.Anything).Return(nil, assert.AnError)

	session := NewSession(TestProjectName, TestAccountName, TestAPIKey)

	resource, err := session.At("people/123").ReadOne()

	if assert.Nil(t, resource) {
		assert.Equal(t, assert.AnError, err)
	}

}

func TestRequest_ReadOne_StretchrError(t *testing.T) {

	mockedTransporter := new(api.MockedTransporter)
	api.ActiveLiveTransporter = mockedTransporter

	// make a response
	response := NewTestResponse(500, nil, []map[string]interface{}{map[string]interface{}{"~message": "Something went wrong"}}, "", nil)
	mockedTransporter.On("MakeRequest", mock.Anything).Return(response, nil)

	session := NewSession(TestProjectName, TestAccountName, TestAPIKey)

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

	responseData := map[string]interface{}{"~count": 2, "~data": []interface{}{map[string]interface{}{"name": "Mat", common.DataFieldID: "ABC"},
		map[string]interface{}{"name": "Tyler", common.DataFieldID: "DEF"}}}

	response := NewTestResponse(200, responseData, nil, "", nil)
	mockedTransporter.On("MakeRequest", mock.Anything).Return(response, nil)

	session := NewSession(TestProjectName, TestAccountName, TestAPIKey)

	resourceCollection, err := session.At("people").ReadMany()

	if assert.NoError(t, err) {
		assert.Equal(t, 2, len(resourceCollection.Resources))
	}

	mockedTransporter.AssertExpectations(t)
	request := mockedTransporter.Calls[0].Arguments[0].(*api.Request)

	assert.Equal(t, request.HTTPMethod(), common.HTTPMethodGet)
	assert.Equal(t, request.Path(), "people")
	assert.Equal(t, request.Body(), []byte(""))

	resource1 := resourceCollection.Resources[0]
	resource2 := resourceCollection.Resources[1]

	assert.Equal(t, resource1.ResourceData()["name"], response.BodyObject().Data().(map[string]interface{})["~data"].([]interface{})[0].(map[string]interface{})["name"])
	assert.Equal(t, resource2.ResourceData()["name"], response.BodyObject().Data().(map[string]interface{})["~data"].([]interface{})[1].(map[string]interface{})["name"])
	assert.Equal(t, resource1.ResourcePath(), "people/ABC")
	assert.Equal(t, resource2.ResourcePath(), "people/DEF")
	assert.Equal(t, resource1.ResourcePath(), "people/ABC")
	assert.Equal(t, resource2.ResourcePath(), "people/DEF")

}

func TestRequest_ReadMany_WithTotal(t *testing.T) {

	mockedTransporter := new(api.MockedTransporter)
	api.ActiveLiveTransporter = mockedTransporter

	// make a response

	responseData := map[string]interface{}{"~count": 2, common.ResponseObjectFieldTotal: 500, "~data": []interface{}{map[string]interface{}{"name": "Mat", common.DataFieldID: "ABC"},
		map[string]interface{}{"name": "Tyler", common.DataFieldID: "DEF"}}}

	response := NewTestResponse(200, responseData, nil, "", nil)
	mockedTransporter.On("MakeRequest", mock.Anything).Return(response, nil)

	session := NewSession(TestProjectName, TestAccountName, TestAPIKey)

	resourceCollection, err := session.At("people").ReadMany()

	if assert.NoError(t, err) {
		assert.Equal(t, 2, len(resourceCollection.Resources))
	}

	mockedTransporter.AssertExpectations(t)
	request := mockedTransporter.Calls[0].Arguments[0].(*api.Request)

	assert.Equal(t, request.HTTPMethod(), common.HTTPMethodGet)
	assert.Equal(t, request.Path(), "people")
	assert.Equal(t, request.Body(), []byte(""))

	assert.Equal(t, resourceCollection.Total, 500)

	resource1 := resourceCollection.Resources[0]
	resource2 := resourceCollection.Resources[1]

	assert.Equal(t, resource1.ResourceData()["name"], response.BodyObject().Data().(map[string]interface{})["~data"].([]interface{})[0].(map[string]interface{})["name"])
	assert.Equal(t, resource2.ResourceData()["name"], response.BodyObject().Data().(map[string]interface{})["~data"].([]interface{})[1].(map[string]interface{})["name"])
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

	session := NewSession(TestProjectName, TestAccountName, TestAPIKey)

	resource, err := session.At("people/123").ReadMany()

	if assert.Nil(t, resource) {
		assert.Equal(t, assert.AnError, err)
	}

}

func TestRequest_ReadMany_StretchrError(t *testing.T) {

	mockedTransporter := new(api.MockedTransporter)
	api.ActiveLiveTransporter = mockedTransporter

	// make a response
	response := NewTestResponse(500, nil, []map[string]interface{}{map[string]interface{}{"~message": "Something went wrong"}}, "", nil)
	mockedTransporter.On("MakeRequest", mock.Anything).Return(response, nil)

	session := NewSession(TestProjectName, TestAccountName, TestAPIKey)

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
	response := NewTestResponse(200, nil, nil, "", api.ChangeInfo(map[string]interface{}{common.ChangeInfoPublicFieldDeleted: 5}))
	mockedTransporter.On("MakeRequest", mock.Anything).Return(response, nil)

	session := NewSession(TestProjectName, TestAccountName, TestAPIKey)

	changeInfo, err := session.At("people/123").Delete()

	if assert.NoError(t, err) {
		assert.NotNil(t, changeInfo)
	}

	mockedTransporter.AssertExpectations(t)
	request := mockedTransporter.Calls[0].Arguments[0].(*api.Request)

	assert.Equal(t, request.HTTPMethod(), common.HTTPMethodDelete)
	assert.Equal(t, request.Path(), "people/123")
	assert.Equal(t, request.Body(), []byte(""))

	assert.Equal(t, changeInfo.Deleted(), 5)

}

func TestRequest_Delete_ReadError(t *testing.T) {

	mockedTransporter := new(api.MockedTransporter)
	api.ActiveLiveTransporter = mockedTransporter

	// make a response
	mockedTransporter.On("MakeRequest", mock.Anything).Return(nil, assert.AnError)

	session := NewSession(TestProjectName, TestAccountName, TestAPIKey)

	resource, err := session.At("people/123").Delete()

	if assert.Nil(t, resource) {
		assert.Equal(t, assert.AnError, err)
	}

}

func TestRequest_Delete_StretchrError(t *testing.T) {

	mockedTransporter := new(api.MockedTransporter)
	api.ActiveLiveTransporter = mockedTransporter

	// make a response
	response := NewTestResponse(500, nil, []map[string]interface{}{map[string]interface{}{"~message": "Something went wrong"}}, "", nil)
	mockedTransporter.On("MakeRequest", mock.Anything).Return(response, nil)

	session := NewSession(TestProjectName, TestAccountName, TestAPIKey)

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
	response := NewTestResponse(200, nil, nil, "", api.ChangeInfo(map[string]interface{}{common.ChangeInfoPublicFieldCreated: 1, common.ChangeInfoPublicFieldDeltas: []interface{}{map[string]interface{}{common.DataFieldID: "hello", common.ChangeInfoPublicFieldDeltaCreated: 123}}}))
	mockedTransporter.On("MakeRequest", mock.Anything).Return(response, nil)

	session := NewSession(TestProjectName, TestAccountName, TestAPIKey)

	changeInfo, err := session.At("people").Create(resource)

	if assert.NoError(t, err) {
		if assert.NotNil(t, changeInfo) {
			mockedTransporter.AssertExpectations(t)
			request := mockedTransporter.Calls[0].Arguments[0].(*api.Request)

			assert.Equal(t, request.HTTPMethod(), common.HTTPMethodPost)
			assert.Equal(t, request.Path(), "people")
			assert.Equal(t, changeInfo.Created(), 1)
			assert.Equal(t, resource.ID(), "hello")
			assert.Equal(t, resource.ResourceData()[common.ChangeInfoPublicFieldDeltaCreated].(float64), 123)
		}
	}

}

func TestRequest_CreateMany(t *testing.T) {

	var resources []*Resource

	resource := MakeResourceAt("people")
	resource.Set("name", "Mat").Set("age", 29)
	resources = append(resources, resource)
	resource = MakeResourceAt("people")
	resource.Set("name", "Tyler").Set("age", 28)
	resources = append(resources, resource)
	resource = MakeResourceAt("people")
	resource.Set("name", "Stacey").Set("age", 29)
	resources = append(resources, resource)

	resourceCollection := NewResourceCollection(resources)

	mockedTransporter := new(api.MockedTransporter)
	api.ActiveLiveTransporter = mockedTransporter

	// make a response
	response := NewTestResponse(200, nil, nil, "", api.ChangeInfo(map[string]interface{}{common.ChangeInfoPublicFieldCreated: 3, common.ChangeInfoPublicFieldDeltas: []interface{}{
		map[string]interface{}{common.DataFieldID: "hello", common.ChangeInfoPublicFieldDeltaCreated: 123},
		map[string]interface{}{common.DataFieldID: "goodbye", common.ChangeInfoPublicFieldDeltaCreated: 456},
		map[string]interface{}{common.DataFieldID: "greetings", common.ChangeInfoPublicFieldDeltaCreated: 789}}}))
	mockedTransporter.On("MakeRequest", mock.Anything).Return(response, nil)

	session := NewSession(TestProjectName, TestAccountName, TestAPIKey)

	changeInfo, err := session.At("people").CreateMany(resourceCollection)

	if assert.NoError(t, err) {
		if assert.NotNil(t, changeInfo) {
			mockedTransporter.AssertExpectations(t)
			request := mockedTransporter.Calls[0].Arguments[0].(*api.Request)

			assert.Equal(t, request.HTTPMethod(), common.HTTPMethodPost)
			assert.Equal(t, request.Path(), "people")
			assert.Equal(t, changeInfo.Created(), 3)

			assert.Equal(t, resourceCollection.Resources[0].ID(), "hello")
			assert.Equal(t, resourceCollection.Resources[0].ResourceData()[common.ChangeInfoPublicFieldDeltaCreated].(float64), 123)
			assert.Equal(t, resourceCollection.Resources[1].ID(), "goodbye")
			assert.Equal(t, resourceCollection.Resources[1].ResourceData()[common.ChangeInfoPublicFieldDeltaCreated].(float64), 456)
			assert.Equal(t, resourceCollection.Resources[2].ID(), "greetings")
			assert.Equal(t, resourceCollection.Resources[2].ResourceData()[common.ChangeInfoPublicFieldDeltaCreated].(float64), 789)
		}
	}

}

func TestRequest_Update_Create(t *testing.T) {

	resource := MakeResourceAt("people")
	resource.Set("name", "Mat").Set("age", 29)

	mockedTransporter := new(api.MockedTransporter)
	api.ActiveLiveTransporter = mockedTransporter

	// make a response
	response := NewTestResponse(200, nil, nil, "", api.ChangeInfo(map[string]interface{}{common.ChangeInfoPublicFieldUpdated: 1, common.ChangeInfoPublicFieldDeltas: []interface{}{map[string]interface{}{common.DataFieldID: "hello", common.ChangeInfoPublicFieldDeltaCreated: 123}}}))
	mockedTransporter.On("MakeRequest", mock.Anything).Return(response, nil)

	session := NewSession(TestProjectName, TestAccountName, TestAPIKey)

	changeInfo, err := session.At("people").Update(resource)

	if assert.NoError(t, err) {
		if assert.NotNil(t, changeInfo) {
			mockedTransporter.AssertExpectations(t)
			request := mockedTransporter.Calls[0].Arguments[0].(*api.Request)

			assert.Equal(t, request.HTTPMethod(), common.HTTPMethodPatch)
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
	response := NewTestResponse(200, nil, nil, "", api.ChangeInfo(map[string]interface{}{common.ChangeInfoPublicFieldCreated: 1, common.ChangeInfoPublicFieldDeltas: []interface{}{map[string]interface{}{common.DataFieldID: "new"}}}))
	mockedTransporter.On("MakeRequest", mock.Anything).Return(response, nil)

	session := NewSession(TestProjectName, TestAccountName, TestAPIKey)

	changeInfo, err := session.At("people").Replace(resource)

	if assert.NoError(t, err) {
		if assert.NotNil(t, changeInfo) {
			mockedTransporter.AssertExpectations(t)
			request := mockedTransporter.Calls[0].Arguments[0].(*api.Request)

			assert.Equal(t, request.HTTPMethod(), common.HTTPMethodPut)
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
	response := NewTestResponse(200, nil, nil, "", api.ChangeInfo(map[string]interface{}{common.ChangeInfoPublicFieldCreated: 1, common.ChangeInfoPublicFieldDeltas: []interface{}{map[string]interface{}{common.DataFieldID: "hello", common.ChangeInfoPublicFieldDeltaCreated: 123}}}))
	mockedTransporter.On("MakeRequest", mock.Anything).Return(response, nil)

	session := NewSession(TestProjectName, TestAccountName, TestAPIKey)

	changeInfo, err := session.At("people").Replace(resource)

	if assert.NoError(t, err) {
		if assert.NotNil(t, changeInfo) {
			mockedTransporter.AssertExpectations(t)
			request := mockedTransporter.Calls[0].Arguments[0].(*api.Request)

			assert.Equal(t, request.HTTPMethod(), common.HTTPMethodPut)
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
	response := NewTestResponse(200, nil, nil, "", api.ChangeInfo(map[string]interface{}{common.ChangeInfoPublicFieldCreated: 1, common.ChangeInfoPublicFieldDeltas: []interface{}{map[string]interface{}{common.DataFieldID: "hello", common.ChangeInfoPublicFieldDeltaCreated: 123}}}))
	mockedTransporter.On("MakeRequest", mock.Anything).Return(response, nil)

	session := NewSession(TestProjectName, TestAccountName, TestAPIKey)

	changeInfo, err := session.At("people").Update(resource)

	if assert.NoError(t, err) {
		if assert.NotNil(t, changeInfo) {
			mockedTransporter.AssertExpectations(t)
			request := mockedTransporter.Calls[0].Arguments[0].(*api.Request)

			assert.Equal(t, request.HTTPMethod(), common.HTTPMethodPatch)
			assert.Equal(t, request.Path(), "people")
			assert.Equal(t, changeInfo.Created(), 1)
		}
	}

}
