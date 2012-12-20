package stretchr

import (
	"net/http"
)

// Repsonse represents a response from a Stretchr request.
type Response struct {
	HttpResponse *http.Response
}

func newResponse(httpResponse *http.Response) *Response {

	repsonse := new(Response)

	// allow the user to have access to the real underlying HTTP response
	repsonse.HttpResponse = httpResponse

	return repsonse
}
