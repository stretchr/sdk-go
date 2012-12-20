package stretchr

import (
	"net/url"
)

// Request represents a kind of interaction that can be made against
// the Stretchr services.
type Request struct {
	session     *Session
	path        string
	queryValues url.Values
	httpMethod  string
}

// NewRequest makes a new Request object, and sets the Session and path.
func NewRequest(session *Session, path string) *Request {
	r := new(Request)
	r.path = path
	r.session = session
	r.queryValues = make(url.Values)
	return r
}

// Path gets the path for this request.
func (r *Request) Path() string {
	return r.path
}

// Session gets the Session object for this request.
func (r *Request) Session() *Session {
	return r.session
}

// QueryValues gets the query values for this request.
func (r *Request) QueryValues() url.Values {
	return r.queryValues
}

// URL generates the absolute URL that will be used to make
// this request.
func (r *Request) URL() (*url.URL, error) {

	urlString := MergeStrings(r.session.Host(), pathSeparator, r.path)

	theUrl, urlErr := url.Parse(urlString)

	if urlErr != nil {
		return nil, urlErr
	}

	// set the query values
	theUrl.RawQuery = r.queryValues.Encode()

	return theUrl, nil

}

/*
	Filtering
*/

// Where adds a filter to the request.
func (r *Request) Where(field, match string) *Request {
	r.queryValues.Add(MergeStrings(filterFieldPrefix, field), match)
	return r
}

/*
	Actions
*/

// Read executes the Request and returns the Response, or an error
// if something went wrong communicating with Stretchr.
func (r *Request) Read() (*Response, error) {

	// set the HTTP method
	r.httpMethod = HttpMethodGet

	// get the transporter to do the work
	return r.session.transporter.MakeRequest(r)
}
