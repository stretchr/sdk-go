package api

import (
	"bytes"
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

	bodyObj := response.BodyObject()

	assert.Equal(t, 200, bodyObj["s"])
	assert.Equal(t, "Mat", bodyObj["d"].(map[string]interface{})["name"])

}
