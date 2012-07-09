package stretchr

import (
	"net/http"
	"testing"
)

func TestSession_Delete(t *testing.T) {

	// use the test requester
	ActiveRequester = ActiveTestRequester.Reset()

	ActiveTestRequester.ResponseToReturn = MakeTestResponseWithData(http.StatusOK, makeStandardResponseObject(http.StatusOK, nil))

	// make a resource
	err := TestSession.Delete("people", "ABC")

	if err != nil {
		t.Errorf("Shouldn't throw error: %s", err)
	}

	AssertLastRequest(t, DeleteMethod, TestSession.Url("people/ABC"), "")

}
