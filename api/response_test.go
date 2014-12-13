package api

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/sdk-go/common"
	"github.com/stretchr/testify/assert"
)

func TestResponse_NewResponse(t *testing.T) {

	response, _ := MakeTestResponse()

	assert.NotNil(t, response.httpResponse)
	assert.Equal(t, response.session, GetTestSession())

}

func TestResponse_HTTPResponse(t *testing.T) {

	response, _ := MakeTestResponse()

	assert.NotNil(t, response.httpResponse)

}

func TestResponse_Session(t *testing.T) {

	response, _ := MakeTestResponse()

	assert.Equal(t, response.Session(), GetTestSession())

}

func TestResponse_HTTPStatusCode(t *testing.T) {

	session := GetTestSession()

	httpResponse := new(http.Response)
	httpResponse.Body = ioutil.NopCloser(bytes.NewBufferString(`{}`))
	httpResponse.StatusCode = http.StatusNotAcceptable
	httpResponse.Header = make(map[string][]string)

	response, err := NewResponse(session, httpResponse)

	if assert.Nil(t, err) {
		assert.Equal(t, http.StatusNotAcceptable, response.HTTPStatusCode())
	}

}

func TestResponse_BodyObject(t *testing.T) {

	response, _ := MakeTestResponseWithBody(`{"~status":200, "~data":{"name":"Mat"}}`)

	bodyObj := response.BodyObject()

	assert.Equal(t, 200, bodyObj["~status"])
	assert.Equal(t, "Mat", bodyObj[common.ResponseObjectFieldData].(map[string]interface{})["name"])

}
