package stretchr

import (
	"net/http"
	"testing"
)

func TestSession_Read(t *testing.T) {

	// use the test requester
	ActiveRequester = ActiveTestRequester.Reset()

	responseData := map[string]interface{}{"name": "Mat", "age": 29, "developer": true, IDKey: "ABC"}
	ActiveTestRequester.ResponseToReturn = MakeTestResponseWithData(http.StatusOK, makeStandardResponseObject(http.StatusOK, responseData))

	// make a resource
	resource, err := TestSession.Read("people", "ABC")

	if err != nil {
		t.Errorf("Shouldn't throw error: %s", err)
	}

	AssertLastRequest(t, ReadMethod, TestSession.Url("people/ABC"), "", "PRIVATE")

	if resource == nil {
		t.Errorf("Resource shouldn't be nil")
	} else {

		AssertEqual(t, "Mat", resource.Get("name"))
		AssertEqual(t, float64(29), resource.Get("age"))
		AssertEqual(t, "ABC", resource.GetID())

	}

}
