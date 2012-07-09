package stretchr

/*
	NOTE: this test file is testing code in http.go
*/

import (
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

/*
	TestRequester

	Replacement object used to test http behaviours without actually
	making HTTP requests.
*/
type TestRequester struct {
	LastMethod     string
	LastURL        string
	LastBody       string
	LastPrivateKey string

	ErrorToReturn    error
	ResponseToReturn *http.Response
}

// MakeRequest stores the parameters and returns the specified ResponseToReturn and ResponseToReturn.
func (r *TestRequester) MakeRequest(method, fullUrl, body, privateKey string) (*StandardResponseObject, *http.Response, error) {

	// save the response bits
	r.LastMethod = method
	r.LastURL = fullUrl
	r.LastBody = body
	r.LastPrivateKey = privateKey

	var sro *StandardResponseObject
	if r.ResponseToReturn != nil {
		sro, _ = ExtractStandardResponseObject(r.ResponseToReturn)
	}

	return sro, r.ResponseToReturn, r.ErrorToReturn
}

func (r *TestRequester) SetClient(c client) {

}

func (r *TestRequester) Client() client {
	return nil
}

func (r *TestRequester) Reset() *TestRequester {
	r.LastMethod = "(none)"
	r.LastURL = "(none)"
	r.LastBody = "(none)"
	r.LastPrivateKey = "(none)"
	r.ErrorToReturn = nil
	r.ResponseToReturn = nil
	return r
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

func MakeOKTestResponse() *http.Response {
	return MakeTestResponse(200, "{\"s\":200,\"w\":true}")
}

// MakeTestResponseWithData makes a test http.Response object initialised with the given
// properties, with the data object as JSON in the response.
func MakeTestResponseWithData(statusCode int, data map[string]interface{}) *http.Response {
	resp := new(http.Response)
	resp.StatusCode = statusCode
	resp.Body = ioutil.NopCloser(strings.NewReader(ToTestJson(data)))
	return resp
}

/*
	TestClient
*/
type TestClient struct {
	LastRequest *http.Request

	ResponseToReturn *http.Response
	ErrorToReturn    error
}

func (c *TestClient) Do(req *http.Request) (*http.Response, error) {
	c.LastRequest = req
	return c.ResponseToReturn, c.ErrorToReturn
}

/*
	DefaultRequester
*/

func TestDefaultRequester_Client(t *testing.T) {

	r := new(DefaultRequester)

	client := new(http.Client)

	r.setClient(client)

	AssertEqual(t, client, r.client())

}

func TestDefaultRequester_Default_Client(t *testing.T) {

	r := new(DefaultRequester)

	client := new(http.Client)

	AssertEqual(t, reflect.TypeOf(client), reflect.TypeOf(r.client()))

}

func TestDefaultRequester_Signing(t *testing.T) {

	r := new(DefaultRequester)

	// use test client
	c := new(TestClient)
	c.ResponseToReturn = MakeOKTestResponse()
	r.setClient(c)

	r.MakeRequest(GetMethod, "http://test.stretchr.com/api/v1/people/ABC?~key=ABC123", "", "PRIVATE")

	url := c.LastRequest.URL.String()
	AssertEqual(t, "http://test.stretchr.com/api/v1/people/ABC?~key=ABC123&~sign=3a4b67d6cba0caab8a5dc13c5dd07dbe11f316ba", url)

}
