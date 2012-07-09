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

func TestworkedFromStatusCode(t *testing.T) {

	AssertEqual(t, true, workedFromStatusCode(200))
	AssertEqual(t, true, workedFromStatusCode(300))
	AssertEqual(t, false, workedFromStatusCode(400))
	AssertEqual(t, false, workedFromStatusCode(500))

}
