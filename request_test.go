package stretchr

import (
	"github.com/stretchrcom/testify/assert"
	"net/url"
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

func TestRequest_Session(t *testing.T) {

	r := new(Request)
	r.session = getTestSession()
	assert.Equal(t, testSession, r.Session())

}

func TestRequest_Path(t *testing.T) {

	r := new(Request)
	r.path = "path"
	assert.Equal(t, "path", r.Path())

}

func TestRequest_QueryValues(t *testing.T) {

	r := new(Request)
	values := make(url.Values)
	r.queryValues = values
	assert.Equal(t, values, r.QueryValues())

}

func TestRequest_Where(t *testing.T) {

	r := NewRequest(getTestSession(), "people")

	returnOfWhere := r.Where("age", "18")

	assert.Equal(t, returnOfWhere, r, ".Where should chain")
	assert.Equal(t, "18", r.queryValues["age"][0])

	// add another where
	r.Where("age", "<30")

	assert.Equal(t, "18", r.queryValues["age"][0])
	assert.Equal(t, "<30", r.queryValues["age"][1])

}

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
