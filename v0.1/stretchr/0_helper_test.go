package stretchr

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func AssertNil(t *testing.T, obj interface{}, message ...string) bool {

	if obj != nil {
		t.Errorf("%s Not nil: %s", message, obj)
		return false
	}
	return true

}

func AssertEqual(t *testing.T, a, b interface{}, message ...string) bool {

	if a != b {
		t.Errorf("%s Not equal. %s != %s.", message, a, b)
		return false
	}
	return true

}

func AssertNotEqual(t *testing.T, a, b interface{}, message ...string) bool {

	if a == b {
		t.Errorf("%s Should not be equal. %s == %s.", message, a, b)
		return false
	}
	return true

}

func AssertOK(t *testing.T, obj interface{}, message ...string) bool {

	if obj == nil {
		t.Errorf("%s Expected not to be nil.", message)
		return false
	}

	return true

}

func AssertContains(t *testing.T, s, contains string, message ...string) bool {

	if !strings.Contains(s, contains) {
		t.Errorf("%s '%s' does not contain '%s'", message, s, contains)
		return false
	}

	return true

}

func AssertNotContains(t *testing.T, s, contains string, message ...string) bool {

	if strings.Contains(s, contains) {
		t.Errorf("%s '%s' should not contain '%s'", message, s, contains)
		return false
	}

	return true

}

func AssertHeader(t *testing.T, request *http.Request, name, value string, message ...string) bool {
	return AssertEqual(t, request.Header.Get(name), value, fmt.Sprintf("%s Header %s should be '%s' but was '%s'", message, name, value, request.Header.Get(name)))
}
