package stretchr

import (
	"fmt"
	"testing"
)

func AssertLastRequest(t *testing.T, method, url, body, privateKey string) {

	AssertEqual(t, method, ActiveTestRequester.LastMethod, "HTTP Method")
	AssertEqual(t, url, ActiveTestRequester.LastURL, "URL")
	AssertEqual(t, body, ActiveTestRequester.LastBody, "Body")
	AssertEqual(t, privateKey, ActiveTestRequester.LastPrivateKey, "Private Key")

}

func ToTestJson(obj map[string]interface{}) string {
	s, err := toJson(obj)

	if err != nil {
		return fmt.Sprintf("{\"Error in ToTestJson\":\"%s\"}", err)
	}

	return s
}
