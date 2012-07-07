package stretchr

import (
	"net/http"
	"strings"
)

// Action represents the object that will be used to describe an action
// to be carried out against the Stretchr services.
type Action struct {

	// Method is the HTTP method that will be used to perform this action
	Method string

	// Path is the 
	Path string

	// Data is a string containing the data for this action.
	Data string

	// Session is a reference to the 
	Session *Session
}

// MakeAction makes a new Action object.
func (s *Session) MakeAction(method, path string) *Action {

	q := new(Action)

	q.Session = s
	q.Method = method
	q.Path = path

	return q

}

// GetRequest gets the http.Request that will actually perform this action.
func (a *Action) GetRequest() (*http.Request, error) {
	return http.NewRequest(a.Method, a.Session.Url(a.Path), strings.NewReader(a.Data))
}

// WithData specifies the data for this action.
// Returns the same action for chaining.
func (a *Action) WithData(data string) *Action {
	a.Data = data
	return a
}
