package stretchr

import (
	"net/url"
	"strconv"
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

// signedUrl generates the absolute URL that will be used to make
// this request.
func (r *Request) signedUrl() (*url.URL, error) {

	urlString := MergeStrings(r.session.host(), pathSeparator, r.path)

	theUrl, urlErr := url.Parse(urlString)

	if urlErr != nil {
		return nil, urlErr
	}

	// set the query values
	theUrl.RawQuery = r.queryValues.Encode()

	// TODO: add security

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

// Limit sets a limit on the number of resources to get back from Stretchr.
func (r *Request) Limit(value int64) *Request {
	r.queryValues.Set(modifierLimit, strconv.FormatInt(value, 10))
	return r
}
