package stretchr

import (
	"testing"
)

func TestMethod(t *testing.T) {

	AssertEqual(t, "GET", ReadMethod)
	AssertEqual(t, "POST", CreateMethod)
	AssertEqual(t, "PUT", UpdateMethod)
	AssertEqual(t, "POST", ReplaceMethod)
	AssertEqual(t, "DELETE", DeleteMethod)

}

func TestWorkedFromStatusCode(t *testing.T) {

	AssertEqual(t, true, WorkedFromStatusCode(200))
	AssertEqual(t, true, WorkedFromStatusCode(300))
	AssertEqual(t, false, WorkedFromStatusCode(400))
	AssertEqual(t, false, WorkedFromStatusCode(500))

}
