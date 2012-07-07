package stretchr

import (
	"net/http"
	"strings"
	"fmt"
)

const (
	Get string = "GET"
	Put string = "PUT"
	Post string = "POST"
	Delete string = "DELETE"
)

func RequestForRead(s Session, path, id string) (*http.Request, error) {
	return http.NewRequest(Get, s.Url(fmt.Sprintf("%s/%s", path, id)), nil)
}

func RequestForReadMany(s Session, path string) (*http.Request, error) {
	return http.NewRequest(Get, s.Url(path), nil)
}

func RequestForCreate(s Session, path, data string) (*http.Request, error) {
	return http.NewRequest(Post, s.Url(path), strings.NewReader(data))
}

func RequestForUpdate(s Session, path, id, data string) (*http.Request, error) {
	return http.NewRequest(Put, s.Url(fmt.Sprintf("%s/%s", path, id)), strings.NewReader(data))
}

func RequestForReplace(s Session, path, id, data string) (*http.Request, error) {
	return http.NewRequest(Post, s.Url(fmt.Sprintf("%s/%s", path, id)), strings.NewReader(data))
}

func RequestForDelete(s Session, path, id string) (*http.Request, error) {
	return http.NewRequest(Delete, s.Url(fmt.Sprintf("%s/%s", path, id)), nil)
}

func RequestForDeleteMany(s Session, path string) (*http.Request, error) {
	return http.NewRequest(Delete, s.Url(path), nil)
}