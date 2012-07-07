package stretchr

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Action represents the object that will be used to describe an action
// to be carried out against the Stretchr services.
//
// Normally, clients will not use this object directly.  Instead they will
// use the CRUD operations on the Session object.
type Action struct {

	// Method is the HTTP method that will be used to perform this action
	Method string

	// Path is the 
	Path string

	// Data is a string containing the data for this action.
	Data string

	// Session is a reference to the 
	Session *Session

	// Values represent the query values that will be made as part of
	// this action.
	Values url.Values
}

// MakeAction makes a new Action object.
func (s *Session) MakeAction(method, path string) *Action {

	q := new(Action)

	q.Session = s
	q.Method = method
	q.Path = path
	q.Values = make(url.Values)

	return q

}

// GetRequest gets the http.Request that will actually perform this action.
//
// Normally clients will not call this method directly, but it is exposed so as
// not to restrict advanced users.
func (a *Action) GetRequest() (*http.Request, error) {
	return http.NewRequest(a.Method, a.Session.Url(a.Path), strings.NewReader(a.Data))
}

// WithData specifies the data for this action.
// Returns the same action for chaining.
func (a *Action) WithData(data string) *Action {
	a.Data = data
	return a
}

/*
	Parametesr
*/

// Limit sets the number of resources that will be returned in this action.
//
// This function returns the same Action for chaining.
func (a *Action) Limit(limit int) *Action {
	a.Values.Set(limitKey, fmt.Sprintf("%d", limit))
	return a
}

// Skip sets the number of resources that will be ignored before they start
// being returned.  Useful for paging.
//
// This function returns the same Action for chaining.
func (a *Action) Skip(skip int) *Action {
	a.Values.Set(skipKey, fmt.Sprintf("%d", skip))
	return a
}

// Order specifies a comma separated list of fields by which to order the results.
//
// This function returns the same Action for chaining.
func (a *Action) Order(order string) *Action {
	a.Values.Set(orderKey, order)
	return a
}

// Context sets a string that will be returned in the response.  Useful for
// lining up responses with requests where context otherwise could be lost.
//
// This function returns the same Action for chaining.
func (a *Action) Context(context string) *Action {
	a.Values.Set(contextKey, context)
	return a
}
