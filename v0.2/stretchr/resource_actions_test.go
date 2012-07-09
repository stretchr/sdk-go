package stretchr

import (
	"net/http"
	"testing"
)

func TestCreate(t *testing.T) {

	// use the test requester
	ActiveRequester = ActiveTestRequester

	responseData := map[string]interface{}{"~id": "ID_RETURNED_BY_SERVER"}
	ActiveTestRequester.ResponseToReturn = MakeTestResponseWithData(http.StatusCreated, makeStandardResponseObject(http.StatusCreated, responseData))

	// make a resource
	r := MakeResource(TestSession, "people")

	// set some data
	r.Set("name", "Mat").Set("age", 29).Set("developer", true)

	err := r.Create()

	if err != nil {
		t.Errorf("Shouldn't throw error: %s", err)
	}

	AssertLastRequest(t, CreateMethod, TestSession.Url("people"), "{\"age\":29,\"developer\":true,\"name\":\"Mat\"}")
	AssertEqual(t, "ID_RETURNED_BY_SERVER", r.Get("~id"), "ID should be set on the resource")

}
