package stretchr

import (
	"testing"
)

func TestMethod(t *testing.T) {

	AssertEqual(t, "GET", GetMethod)
	AssertEqual(t, "PUT", PutMethod)
	AssertEqual(t, "POST", PostMethod)
	AssertEqual(t, "DELETE", DeleteMethod)

	AssertEqual(t, GetMethod, ReadMethod)
	AssertEqual(t, PostMethod, CreateMethod)
	AssertEqual(t, PutMethod, UpdateMethod)
	AssertEqual(t, PostMethod, ReplaceMethod)
	AssertEqual(t, DeleteMethod, DeleteMethod)

}

func TestworkedFromStatusCode(t *testing.T) {

	AssertEqual(t, true, workedFromStatusCode(200))
	AssertEqual(t, true, workedFromStatusCode(300))
	AssertEqual(t, false, workedFromStatusCode(400))
	AssertEqual(t, false, workedFromStatusCode(500))

}
