package stretchr

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const (
	// NoBody is a string that represents no body in a request.
	//
	// If the body is this value, no body will be included
	// (as opposed to a body with an empty string). 
	NoBody string = ""
)

// Requester is an interface describing objects capable of making and processing
// HTTP requests.
type Requester interface {

	// MakeRequest makes a request and returns the response.
	MakeRequest(method, fullUrl, body, publicKey, privateKey string) (*StandardResponseObject, *http.Response, error)
}

// DefaultRequester is a Requester object that makes real HTTP requests.
type DefaultRequester struct {
	// client holds the object by which HTTP requests will be made.
	_client client
}

// MakeRequest makes a request and returns the response.
func (r *DefaultRequester) MakeRequest(method, fullUrl, body, publicKey, privateKey string) (*StandardResponseObject, *http.Response, error) {

	// get the client we'll use
	client := r.client()

	// add the public key
	if strings.Contains(fullUrl, "?") {
		fullUrl = fmt.Sprintf("%s&%s=%s", fullUrl, PublicKeyKey, publicKey)
	} else {
		fullUrl = fmt.Sprintf("%s?%s=%s", fullUrl, PublicKeyKey, publicKey)
	}

	// sign the request
	signedUrl, signUrlErr := GetSignedURL(method, fullUrl, body, privateKey)

	if signUrlErr != nil {
		return nil, nil, signUrlErr
	}

	// make the request
	var request *http.Request
	var requestErr error

	if body == NoBody {
		request, requestErr = http.NewRequest(method, signedUrl, nil)
	} else {
		request, requestErr = http.NewRequest(method, signedUrl, strings.NewReader(body))
	}

	if requestErr != nil {
		return nil, nil, requestErr
	}

	// make the request
	response, doErr := client.Do(request)

	if doErr != nil {
		return nil, response, doErr
	}

	// get the StandardResponseObject
	sro, sroErr := ExtractStandardResponseObject(response)

	log.Printf("error: %s", sroErr)

	if sroErr != nil {
		return nil, response, sroErr
	}

	// return the SRO
	return sro, response, nil
}

// SetClient tells the Requester which Client to use to make requests.
func (r *DefaultRequester) setClient(c client) {
	r._client = c
}

// Client gets the Client to use when making requests.
func (r *DefaultRequester) client() client {

	if r._client == nil {
		r._client = new(http.Client)
	}

	return r._client
}

// ActiveRequester is the Requester object this code will use to make requests.
//
// Only change this value if you are writing test code, or if you decide to have more
// control over the actual HTTP requests made by this library.
var ActiveRequester Requester = new(DefaultRequester)
