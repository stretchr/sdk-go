package api

import (
	"github.com/stretchrcom/stretchr-sdk-go/common"
	"github.com/stretchrcom/testify/assert"
	"io/ioutil"
	"testing"
)

var mockedTransporter *MockedTransporter

func getTestSession() *Session {
	testSession = NewSession("test", "123", "456")
	mockedTransporter = new(MockedTransporter)
	testSession.transporter = mockedTransporter
	return testSession
}

func TestNewRequest(t *testing.T) {

	r := NewRequest(getTestSession(), "people")

	assert.Equal(t, "people", r.path, "path")
	assert.Equal(t, GetTestSession(), r.session, "session")
	assert.NotNil(t, r.queryValues, "queryValues")

}

func TestRequest_Session(t *testing.T) {

	r := NewRequest(getTestSession(), "people")

	assert.Equal(t, testSession, r.Session())

}

func TestRequest_signedURL(t *testing.T) {

	r := NewRequest(getTestSession(), "people/123")
	r.Where("field", "match").Where("field2", "match2")

	fullUrl, urlErr := r.signedUrl()

	if assert.Nil(t, urlErr) {
		urlString := fullUrl.String()
		assert.Contains(t, urlString, "http://test.stretchr.com/api/v1/people/123?")
		assert.Contains(t, urlString, "%3Afield=match")
		assert.Contains(t, urlString, "%3Afield2=match2")
		assert.Contains(t, urlString, "~key=123")
		assert.Contains(t, urlString, "~sign=5ce10ef6c27c01d919e5d879bcee4090801422c1")
	}

}

func TestRequest_hasBody(t *testing.T) {

	r := new(Request)

	assert.False(t, r.hasBody())

	r.body = []byte("Hello Stretchr")

	assert.True(t, r.hasBody())

}

func TestRequest_SetBodyObject(t *testing.T) {

	r := NewRequest(getTestSession(), "people/123")

	obj := make(map[string]interface{})
	obj["name"] = "Mat"

	r.setBodyObject(obj)
	expectedBody, _ := r.session.codec.Marshal(obj, nil)
	assert.Equal(t, r.body, expectedBody)

}

func TestRequest_httpRequest(t *testing.T) {

	r := NewRequest(getTestSession(), "people/123")
	r.body = []byte("Hello Stretchr")

	httpRequest, _ := r.httpRequest()

	assert.Equal(t, httpRequest.Method, r.httpMethod)
	expectedUrl, _ := r.signedUrl()
	assert.Equal(t, httpRequest.URL.String(), expectedUrl.String())
	expectedBody, _ := ioutil.ReadAll(httpRequest.Body)
	assert.Equal(t, expectedBody, r.body)

}

func TestRequest_QueryValues(t *testing.T) {

	r := NewRequest(getTestSession(), "people/123")
	r.Where("name", "Mat")

	assert.Equal(t, r.queryValues, r.QueryValues())

}

/*
	Properties
*/
func TestRequest_HttpMethod(t *testing.T) {

	request := new(Request)
	request.httpMethod = "GET"

	assert.Equal(t, request.httpMethod, request.HttpMethod())

}

func TestRequest_Body(t *testing.T) {

	request := new(Request)
	request.body = []byte("Byte")

	assert.Equal(t, request.body, request.Body())

}

func TestRequest_Path(t *testing.T) {

	request := new(Request)
	request.path = "/path"

	assert.Equal(t, request.path, request.Path())

}

/*
	Filtering
*/

func TestRequest_Where(t *testing.T) {

	r := NewRequest(getTestSession(), "people")

	returnOfWhere := r.Where("age", "18")

	assert.Equal(t, returnOfWhere, r, ".Where should chain")

	// add another where
	r.Where("age", "<30")

	if assert.Equal(t, 2, len(r.queryValues[":age"]), "Should be two values for :age") {
		assert.Equal(t, "18", r.queryValues[":age"][0])
		assert.Equal(t, "<30", r.queryValues[":age"][1])
	}

}

func TestRequest_Limit(t *testing.T) {

	r := NewRequest(getTestSession(), "people")

	returnOfLimit := r.Limit(50)

	assert.Equal(t, returnOfLimit, r, ".Limit should chain")

	assert.Equal(t, "50", r.queryValues[common.ModifierLimit][0])

}

func TestRequest_Skip(t *testing.T) {

	r := NewRequest(getTestSession(), "people")

	returnOfSkip := r.Skip(20)

	assert.Equal(t, returnOfSkip, r, ".Skip should chain")

	assert.Equal(t, "20", r.queryValues[common.ModifierSkip][0])

}

func TestRequest_WithParam(t *testing.T) {

	r := NewRequest(getTestSession(), "people")

	assert.Equal(t, r, r.WithParam("name", "Mat"))
	assert.Equal(t, r, r.WithParam("age", "29"))

	assert.Equal(t, "Mat", r.queryValues["name"][0])
	assert.Equal(t, "29", r.queryValues["age"][0])

}
