package stretchr

import (
	"net/http"
	"strings"
	"fmt"
)

const (
	// Get is a constant representing the GET HTTP method
	Get string = "GET"
	
	// Put is a constant representing the PUT HTTP method
	Put string = "PUT"
	
	// Post is a constant representing the POST HTTP method
	Post string = "POST"
	
	// Delete is a constant representing the DELETE HTTP method
	Delete string = "DELETE"
)

// RequestForRead gets the underlying http.Request that will be used to read a resource.
func RequestForRead(s Session, path, id string) (*http.Request, error) {
	return http.NewRequest(Get, s.Url(fmt.Sprintf("%s/%s", path, id)), nil)
}

// RequestForReadMany gets the underlying http.Request that will be used to read a collection of resources.
func RequestForReadMany(s Session, path string) (*http.Request, error) {
	return http.NewRequest(Get, s.Url(path), nil)
}

// RequestForCreate gets the underlying http.Request that will be used to create a resource.
func RequestForCreate(s Session, path, data string) (*http.Request, error) {
	return http.NewRequest(Post, s.Url(path), strings.NewReader(data))
}

// RequestForUpdate gets the underlying http.Request that will be used to update a resource.
func RequestForUpdate(s Session, path, id, data string) (*http.Request, error) {
	return http.NewRequest(Put, s.Url(fmt.Sprintf("%s/%s", path, id)), strings.NewReader(data))
}

// RequestForReplace gets the underlying http.Request that will be used to replace a resource.
func RequestForReplace(s Session, path, id, data string) (*http.Request, error) {
	return http.NewRequest(Post, s.Url(fmt.Sprintf("%s/%s", path, id)), strings.NewReader(data))
}

// RequestForDelete gets the underlying http.Request that will be used to delete a resource.
func RequestForDelete(s Session, path, id string) (*http.Request, error) {
	return http.NewRequest(Delete, s.Url(fmt.Sprintf("%s/%s", path, id)), nil)
}

// RequestForDeleteMany gets the underlying http.Request that will be used to delete resources.
func RequestForDeleteMany(s Session, path string) (*http.Request, error) {
	return http.NewRequest(Delete, s.Url(path), nil)
}