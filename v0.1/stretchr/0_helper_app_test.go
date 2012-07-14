package stretchr

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func testBaseUrl() string {
	return fmt.Sprintf("%s://%s.%s/api/%s/", "http", "test", DefaultHost, DefaultVersion)
}

func testSession() *Session {
	return InProject("test").WithKeys("PUB", "PRIV")
}

func AssertRequest(t *testing.T, request *http.Request, method, url string) bool {
	return AssertEqual(t, method, request.Method, "Method") && AssertEqual(t, fmt.Sprintf("%s%s", testBaseUrl(), url), request.URL.String(), "URL")
}

func AssertNoRequestBody(t *testing.T, request *http.Request) bool {
	b := AssertRequestBody(t, request, "")
	if !b {
		t.Errorf("(Request body expected to be empty)")
	}
	return b
}

func AssertRequestBody(t *testing.T, request *http.Request, body string) bool {

	var actualBody []byte

	if request.Body == nil {
		actualBody = make([]byte, 0)
	} else {

		var err error
		actualBody, err = ioutil.ReadAll(request.Body)

		if err != nil {
			panic(fmt.Sprintf("unable to read the request body: %s", err))
		}

	}

	return AssertEqual(t, body, string(actualBody), "Request Body")

}
