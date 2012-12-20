package stretchr

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

var testSession *Session
var mockedTransporter *MockedTransporter

func getTestSession() *Session {
	testSession = NewSession("test")
	mockedTransporter = new(MockedTransporter)
	testSession.transporter = mockedTransporter
	return testSession
}

func TestNewRequest(t *testing.T) {

	r := NewRequest(getTestSession(), "people")

	assert.Equal(t, "people", r.path, "path")
	assert.Equal(t, testSession, r.session, "session")
	assert.NotNil(t, r.queryValues, "queryValues")

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
	}

}

func TestRequest_SetBodyObject(t *testing.T) {

	r := NewRequest(getTestSession(), "people/123")

	obj := make(map[string]interface{})
	obj["name"] = "Mat"

	r.setBodyObject(obj)
	expectedBody, _ := ObjectToBytes(obj) //#codecs
	assert.Equal(t, r.body, expectedBody)

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

	assert.Equal(t, "50", r.queryValues[modifierLimit][0])

}
