package api

import (
	"github.com/stretchr/sdk-go/common"
	stewstrings "github.com/stretchr/stew/strings"
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

// url generates the absolute URL that will be used to make
// this request.
func (r *Request) url() (*url.URL, error) {

	urlString := stewstrings.MergeStrings(r.session.host(), common.PathSeparator, r.path)

	theUrl, urlErr := url.Parse(urlString)

	if urlErr != nil {
		return nil, urlErr
	}

	// set the query values
	theUrl.RawQuery = r.queryValues.Encode()

	urlString = theUrl.String()

	if strings.Contains(urlString, "?") {
		urlString = stewstrings.MergeStrings(urlString, "&", common.ParameterAPIKey, "=", r.session.apiKey)
	} else {
		urlString = stewstrings.MergeStrings(urlString, "?", common.ParameterAPIKey, "=", r.session.apiKey)
	}

	theUrl, urlErr = url.Parse(urlString)

	if urlErr != nil {
		return nil, urlErr
	}

	return theUrl, nil

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

	requestURL, urlErr := r.url()

	if urlErr != nil {
		return nil, urlErr
	}

	if r.hasBody() {
		httpRequest, requestErr = http.NewRequest(r.httpMethod, requestURL.String(), strings.NewReader(string(r.body)))
	} else {
		httpRequest, requestErr = http.NewRequest(r.httpMethod, requestURL.String(), nil)
	}

	return httpRequest, requestErr

}

/*
	Parameters
*/

// QueryValues gets the url.Values that will be passed in the querystring
// when this request is made.
func (r *Request) QueryValues() url.Values {
	return r.queryValues
}

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

// Skip sets the number of resources to skip before getting them back from Stretchr.
func (r *Request) Skip(value int64) *Request {
	return r.WithParam(common.ModifierSkip, strconv.FormatInt(value, 10))
}

// Page sets the page of resources to get by specifying the appropriate Limit and Skip
// values.
func (r *Request) Page(pageNumber, resourcesPerPage int64) *Request {
	return r.Limit(resourcesPerPage).Skip(resourcesPerPage * (pageNumber - 1))
}
