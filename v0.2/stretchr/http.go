package stretchr

const (
	ReadMethod    string = "GET"
	CreateMethod  string = "POST"
	UpdateMethod  string = "PUT"
	ReplaceMethod string = "POST"
	DeleteMethod  string = "DELETE"
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
