package api

import (
	"github.com/stretchrcom/signature"
	stewstrings "github.com/stretchrcom/stew/strings"
	"github.com/stretchrcom/stretchr-sdk-go/common"
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

	urlString := stewstrings.MergeStrings(r.session.host(), common.PathSeparator, r.path)

	theUrl, urlErr := url.Parse(urlString)

	if urlErr != nil {
		return nil, urlErr
	}

	// set the query values
	theUrl.RawQuery = r.queryValues.Encode()

	signedURLString, signErr := signature.GetSignedURL(r.httpMethod, theUrl.String(), string(r.body), r.session.publicKey, r.session.privateKey)

	if signErr != nil {
		return nil, signErr
	}

	signedURL, signedURLErr := url.Parse(signedURLString)

	if signedURLErr != nil {
		return nil, signedURLErr
	}

	return signedURL, nil

}

// setBodyObject sets the bytes in .body to a marshalled version of the specified
// object.
func (r *Request) setBodyObject(object interface{}) error {
	var err error
	r.body, err = r.session.codec.Marshal(object, nil)
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
	Parameters
*/

// WithParam sets a query parameter in the request.
func (r *Request) WithParam(key, value string) *Request {
	r.queryValues.Add(key, value)
	return r
}

/*
	Properties
*/

// Path gets the path for this request.
func (r *Request) Path() string {
	return r.path
}

// HttpMethod gets the HTTP Method that this request will use.
func (r *Request) HttpMethod() string {
	return r.httpMethod
}

// Body gets the bytes that make up the body of this request.
func (r *Request) Body() []byte {
	return r.body
}

// Session gets the session object that this request will use when being made.
func (r *Request) Session() *Session {
	return r.session
}

/*
	Filtering
*/

// Where adds a filter to the request.
func (r *Request) Where(field, match string) *Request {
	return r.WithParam(stewstrings.MergeStrings(common.FilterFieldPrefix, field), match)
}

// Limit sets a limit on the number of resources to get back from Stretchr.
func (r *Request) Limit(value int64) *Request {
	return r.WithParam(common.ModifierLimit, strconv.FormatInt(value, 10))
}
