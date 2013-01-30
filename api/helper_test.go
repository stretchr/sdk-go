package api

import (
	"bytes"
	"fmt"
	"github.com/stretchrcom/sdk-go/common"
	"github.com/stretchrcom/signature"
	"io/ioutil"
	"net/http"
)

func MakeTestResponse() (*Response, error) {

	res, err := MakeTestResponseWithBody(`{"~s":200}`)

	if err != nil {
		panic(fmt.Sprintf("MakeTestResponse: %s", err))
	}

	return res, nil

}

func MakeTestResponseWithBody(body string) (*Response, error) {

	testHttpResponse := new(http.Response)
	testHttpResponse.Body = ioutil.NopCloser(bytes.NewBufferString(body))
	testHttpResponse.StatusCode = 200
	testHttpResponse.Header = make(map[string][]string)
	testHttpResponse.Header[common.HeaderResponseHash] = []string{signature.HashWithKeys([]byte(body), []byte(testSession.publicKey), []byte(testSession.privateKey))}

	return NewResponse(GetTestSession(), testHttpResponse)

}

var testSession *Session

func GetTestSession() *Session {
	if testSession == nil {
		testSession = NewSession("project", "publicKey", "privateKey")
	}
	return testSession
}

func GetTestHttpResponse() *http.Response {
	testHttpResponse := new(http.Response)
	testHttpResponse.Body = ioutil.NopCloser(bytes.NewBufferString(`{"~s":200}`))
	return testHttpResponse
}
