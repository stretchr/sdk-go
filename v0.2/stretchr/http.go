package stretchr

import (
	"net/http"
)

const (
	GetMethod    string = "GET"
	PostMethod   string = "POST"
	PutMethod    string = "PUT"
	DeleteMethod string = "DELETE"

	ReadMethod    string = GetMethod
	CreateMethod  string = PostMethod
	UpdateMethod  string = PutMethod
	ReplaceMethod string = PostMethod
)

const (
	/*
	   HTTP OK status codes
	*/

	// StatusCodesOKMinimum represents the lowest allowed HTTP status code for requests
	// that are to be considered OK.  Status codes between StatusCodesOKMinimum and StatusCodesOKMaximum
	// will decide the 'Worked' output in the standard response object.
	StatusCodesOKMinimum int = 100

	// StatusCodesOKMaximum represents the highest allowed HTTP status code for requests
	// that are to be considered OK.  Status codes between StatusCodesOKMinimum and StatusCodesOKMaximum
	// will decide the 'Worked' output in the standard response object.
	StatusCodesOKMaximum int = 399
)

// workedFromStatusCode gets whether the request was successful based on the given
// HTTP status code.
func workedFromStatusCode(statusCode int) bool {
	return statusCode >= StatusCodesOKMinimum && statusCode <= StatusCodesOKMaximum
}

// Client describes objects that can perform HTTP requests.
type Client interface {
	// Do sends an HTTP request and returns an HTTP response
	Do(req *http.Request) (*http.Response, error)
}
