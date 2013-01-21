package stretchr

import (
	"github.com/stretchrcom/stretchr-sdk-go/api"
	"github.com/stretchrcom/stretchr-sdk-go/common"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

var testSession *Session
var mockedTransporter *api.MockedTransporter

func getTestSession() *Session {
	testSession = NewSession("test", "123", "456")
	testSession.underlyingSession = api.NewSession("project", "publicKey", "privateKey")
	mockedTransporter = new(api.MockedTransporter)
	testSession.SetTransporter(mockedTransporter)
	return testSession
}

func TestNewRequest(t *testing.T) {

	r := NewRequest(getTestSession(), "people/123")

	if assert.NotNil(t, r) {
		assert.Equal(t, r.session, testSession)
		assert.Equal(t, r.UnderlyingRequest.Session(), testSession.underlyingSession)
		assert.Equal(t, r.UnderlyingRequest.Path(), "people/123")
	}

}

func TestSession(t *testing.T) {

	r := NewRequest(getTestSession(), "people/123")
	assert.Equal(t, r.Session(), testSession)

}

func TestRequest_WithParam(t *testing.T) {

	r := NewRequest(getTestSession(), "people")

	assert.Equal(t, r, r.WithParam("name", "Mat"))
	assert.Equal(t, r, r.WithParam("age", "29"))

	assert.Equal(t, "Mat", r.UnderlyingRequest.QueryValues()["name"][0])
	assert.Equal(t, "29", r.UnderlyingRequest.QueryValues()["age"][0])

}

func TestRequest_Where(t *testing.T) {

	r := NewRequest(getTestSession(), "people/123")
	returnOfWhere := r.Where("age", "18")

	assert.Equal(t, returnOfWhere, r, ".Where should chain")

	// add another where
	r.Where("age", "<30")

	if assert.Equal(t, 2, len(r.UnderlyingRequest.QueryValues()[":age"]), "Should be two values for :age") {
		assert.Equal(t, "18", r.UnderlyingRequest.QueryValues()[":age"][0])
		assert.Equal(t, "<30", r.UnderlyingRequest.QueryValues()[":age"][1])
	}

}

func TestRequest_Limit(t *testing.T) {

	r := NewRequest(getTestSession(), "people/123")

	returnOfLimit := r.Limit(50)

	assert.Equal(t, returnOfLimit, r, ".Limit should chain")

	assert.Equal(t, "50", r.UnderlyingRequest.QueryValues()[common.ModifierLimit][0])

}

func TestRequest_Skip(t *testing.T) {

	r := NewRequest(getTestSession(), "people")

	returnOfSkip := r.Skip(20)

	assert.Equal(t, returnOfSkip, r, ".Skip should chain")

	assert.Equal(t, "20", r.UnderlyingRequest.QueryValues()[common.ModifierSkip][0])

}

func TestRequest_Page(t *testing.T) {

	r := NewRequest(getTestSession(), "people")

	returnOfPage := r.Page(1, 50)
	assert.Equal(t, r, returnOfPage, "Page should chain")

	assert.Equal(t, "0", r.UnderlyingRequest.QueryValues()[common.ModifierSkip][0])
	assert.Equal(t, "50", r.UnderlyingRequest.QueryValues()[common.ModifierLimit][0])

	r = NewRequest(getTestSession(), "people").Page(2, 50)
	assert.Equal(t, "50", r.UnderlyingRequest.QueryValues()[common.ModifierSkip][0])
	assert.Equal(t, "50", r.UnderlyingRequest.QueryValues()[common.ModifierLimit][0])

	r = NewRequest(getTestSession(), "people").Page(3, 50)
	assert.Equal(t, "100", r.UnderlyingRequest.QueryValues()[common.ModifierSkip][0])
	assert.Equal(t, "50", r.UnderlyingRequest.QueryValues()[common.ModifierLimit][0])

}
