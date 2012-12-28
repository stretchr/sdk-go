package api

import (
	"bytes"
	"github.com/stretchrcom/stew/objects"
	"github.com/stretchrcom/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestResponse_NewResponse(t *testing.T) {

	response, _ := MakeTestResponse()

	assert.NotNil(t, response.httpResponse)
	assert.Equal(t, response.session, GetTestSession())

}

func TestResponse_HttpResponse(t *testing.T) {

	response, _ := MakeTestResponse()

	assert.NotNil(t, response.httpResponse)

}

func TestResponse_Session(t *testing.T) {

	response, _ := MakeTestResponse()

	assert.Equal(t, response.Session(), GetTestSession())

}

func TestResponse_HttpStatusCode(t *testing.T) {

	httpResponse := new(http.Response)
	httpResponse.Body = ioutil.NopCloser(bytes.NewBufferString(`{}`))
	httpResponse.StatusCode = http.StatusNotAcceptable
	response, _ := NewResponse(GetTestSession(), httpResponse)

	assert.Equal(t, http.StatusNotAcceptable, response.HttpStatusCode())

}

func TestResponse_BodyObject(t *testing.T) {

	response, _ := MakeTestResponseWithBody(`{"s":200,"d":{"name":"Mat"}}`)

	bodyObj := objects.Map(response.BodyObject().(map[string]interface{}))

	assert.Equal(t, 200, bodyObj.Get("s"))
	assert.Equal(t, "Mat", bodyObj.Get("d.name"))

}

func TestResponse_SingleBodyObject(t *testing.T) {

	response, _ := MakeTestResponseWithBody(`{"s":200,"d":{"name":"Mat"}}`)

	bodyObj := objects.Map(response.SingleBodyObject())

	assert.Equal(t, 200, bodyObj.Get("s"))
	assert.Equal(t, "Mat", bodyObj.Get("d.name"))

}

func TestResponse_SingleBodyObject_InvalidObjectType(t *testing.T) {

	response, _ := MakeTestResponseWithBody(`[{"s":200,"d":{"name":"Mat"}}]`)
	assert.Panics(t, func() {
		response.SingleBodyObject()
	}, "Calling SingleBodyObject on an array should panic")

}

func TestResponse_MultipleBodyObjects(t *testing.T) {

	response, _ := MakeTestResponseWithBody(`[{"s":200,"d":{"name":"Mat"}},{"s":200,"d":{"name":"Mat"}}]`)

	bodyObjs := response.MultipleBodyObjects()

	bodyObj := objects.Map(bodyObjs[0])

	assert.Equal(t, 200, bodyObj.Get("s"))
	assert.Equal(t, "Mat", bodyObj.Get("d.name"))

}

func TestResponse_MultipleBodyObjects_InvalidObjectType(t *testing.T) {

	response, _ := MakeTestResponseWithBody(`{"s":200,"d":{"name":"Mat"}}`)
	assert.Panics(t, func() {
		response.MultipleBodyObjects()
	}, "Calling MultipleBodyObjects on a single object should panic")

}
