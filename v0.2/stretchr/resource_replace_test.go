package stretchr

import (
	"net/http"
	"testing"
)

func TestResource_Replace(t *testing.T) {

	// use the test requester
	ActiveRequester = ActiveTestRequester.Reset()

	responseData := map[string]interface{}{"surname": "Ryer"}
	ActiveTestRequester.ResponseToReturn = MakeTestResponseWithData(http.StatusCreated, makeStandardResponseObject(http.StatusCreated, responseData))

	// make a resource
	r := TestSession.MakeResource("people")

	// set some data
	r.Set(IDKey, "ABC").Set("name", "Mat").Set("age", 29).Set("developer", true)

	err := r.Replace()

	if err != nil {
		t.Errorf("Shouldn't throw error: %s", err)
	}

	AssertLastRequest(t, ReplaceMethod, TestSession.Url("people/ABC"), "{\"age\":29,\"developer\":true,\"name\":\"Mat\",\"~id\":\"ABC\"}")

}
