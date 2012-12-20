package stretchr

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Request represents a kind of interaction that can be made against
// the Stretchr services.
type Request struct {
	session     *Session
	path        string
	queryValues url.Values
	httpMethod  string
	body        []byte
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

// setBodyObject sets the bytes in .body to a marshalled version of the specified
// object.
func (r *Request) setBodyObject(object interface{}) error {
	var err error
	r.body, err = ObjectToBytes(object) // TODO: #codecs
	return err
}

func (r *Request) hasBody() bool {
	return len(r.body) > 0
}

// httpRequest gets the http.Request that will be used to perform
// this request.
func (r *Request) httpRequest() (*http.Request, error) {

	var httpRequest *http.Request
	var requestErr error

	signedUrl, urlErr := r.signedUrl()

	if urlErr != nil {
		return nil, urlErr
	}

	if r.hasBody() {
		httpRequest, requestErr = http.NewRequest(r.httpMethod, signedUrl.String(), strings.NewReader(string(r.body)))
	} else {
		httpRequest, requestErr = http.NewRequest(r.httpMethod, signedUrl.String(), nil)
	}

	return httpRequest, requestErr

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
