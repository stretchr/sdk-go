package stretchr

import (
	"net/http"
	"testing"
)

func TestResource_Delete(t *testing.T) {

	// use the test requester
	ActiveRequester = ActiveTestRequester.Reset()

	ActiveTestRequester.ResponseToReturn = MakeTestResponseWithData(http.StatusOK, makeStandardResponseObject(http.StatusOK, nil))

	// make a resource
	r := TestSession.MakeResource("people").SetID("ABC")
	err := r.Delete()

	if err != nil {
		t.Errorf("Shouldn't throw error: %s", err)
	}

	AssertLastRequest(t, DeleteMethod, TestSession.Url("people/ABC"), "")

}
