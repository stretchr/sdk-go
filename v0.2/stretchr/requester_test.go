package stretchr

/*
	NOTE: this test file is testing code in http.go
*/

import (
	"io/ioutil"
	"net/http"
	"strings"
)

/*
	TestRequester

	Replacement object used to test http behaviours without actually
	making HTTP requests.
*/
type TestRequester struct {
	LastMethod string
	LastURL    string
	LastBody   string

	ErrorToReturn    error
	ResponseToReturn *http.Response
}

// MakeRequest stores the parameters and returns the specified ResponseToReturn and ResponseToReturn.
func (r *TestRequester) MakeRequest(method, fullUrl, body string) (*StandardResponseObject, *http.Response, error) {

	// save the response bits
	r.LastMethod = method
	r.LastURL = fullUrl
	r.LastBody = body

	var sro *StandardResponseObject
	if r.ResponseToReturn != nil {
		sro, _ = ExtractStandardResponseObject(r.ResponseToReturn)
	}

	return sro, r.ResponseToReturn, r.ErrorToReturn
}

// ActiveTestRequester represents an instance of the TestRequester.
var ActiveTestRequester *TestRequester = new(TestRequester)

// MakeTestResponse makes a test http.Response object initialised with the given
// properties.
func MakeTestResponse(statusCode int, body string) *http.Response {
	resp := new(http.Response)
	resp.StatusCode = statusCode
	resp.Body = ioutil.NopCloser(strings.NewReader(body))
	return resp
}

// MakeTestResponseWithData makes a test http.Response object initialised with the given
// properties, with the data object as JSON in the response.
func MakeTestResponseWithData(statusCode int, data map[string]interface{}) *http.Response {
	resp := new(http.Response)
	resp.StatusCode = statusCode
	resp.Body = ioutil.NopCloser(strings.NewReader(ToTestJson(data)))
	return resp
}
