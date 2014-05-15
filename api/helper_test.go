package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func MakeTestResponse() (*Response, error) {

	res, err := MakeTestResponseWithBody(`{"~status":200}`)

	if err != nil {
		panic(fmt.Sprintf("MakeTestResponse: %s", err))
	}

	return res, nil

}

func MakeTestResponseWithBody(body string) (*Response, error) {

	testHTTPResponse := new(http.Response)
	testHTTPResponse.Body = ioutil.NopCloser(bytes.NewBufferString(body))
	testHTTPResponse.StatusCode = 200
	testHTTPResponse.Header = make(map[string][]string)

	return NewResponse(GetTestSession(), testHTTPResponse)

}

var testSession *Session

func GetTestSession() *Session {
	if testSession == nil {
		testSession = NewSession("project", "account", "apiKey")
	}
	return testSession
}

func GettestHTTPResponse() *http.Response {
	testHTTPResponse := new(http.Response)
	testHTTPResponse.Body = ioutil.NopCloser(bytes.NewBufferString(`{"~status":200}`))
	return testHTTPResponse
}
