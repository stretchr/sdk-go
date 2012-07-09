package stretchr

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var UnknownError = errors.New("Something went wrong, not sure what - sorry.")

// StandardResponseObject is the top level container for all responses.
type StandardResponseObject struct {

	// Status code contains the HTTP status code of the response
	StatusCode int

	// Errors contains a collection of errors that occurred during the processing of the requwest
	Errors []interface{}

	// Worked is a quick way to find out if a request was successful or not
	Worked bool

	// Data contains the data for the response
	Data map[string]interface{}

	// Context holds the context modifier value to make it easy to line
	// responses up with requests
	Context string
}

// GetError gets the first error from the Errors array, or returns UnknownError.
//
// Check the Worked field before calling this method since this method
// will always return an error of some sort.
func (sro *StandardResponseObject) GetError() error {

	if len(sro.Errors) > 0 {
		return errors.New(fmt.Sprintf("%s", sro.Errors[0].(map[string]interface{})["Message"]))
	}

	return UnknownError

}

// ExtractStandardResponseObject extracts the StandardResponseObject from the specified
// http.Response.
func ExtractStandardResponseObject(response *http.Response) (*StandardResponseObject, error) {

	obj := new(StandardResponseObject)

	// set the real HTTP method
	obj.StatusCode = response.StatusCode

	// set 'worked'
	obj.Worked = workedFromStatusCode(obj.StatusCode)

	// read the actual response object
	responseString, responseStringErr := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	if responseStringErr != nil {
		return nil, responseStringErr
	}

	// get the object from JSON
	respObj, jsonErr := fromJson(string(responseString))

	if jsonErr != nil {
		return nil, jsonErr
	}

	// set the data if there is some
	if respObj["d"] != nil {
		obj.Data = respObj["d"].(map[string]interface{})
	}

	/*
		// set the errors if there are any
		if respObj["e"] != nil {
			obj.Errors = respObj["e"].([]interface{})
		}
	*/

	return obj, nil

}
