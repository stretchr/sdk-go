package api

import (
	"github.com/stretchr/sdk-go/common"
)

const (
	responseObjectFieldErrorStrings string = "_errors"
)

// ResponseObject is a map[string]interface{} with some nice methods on it to
// make working with a response object easier.
type ResponseObject map[string]interface{}

// StatusCode returns the status code in the response
func (o ResponseObject) StatusCode() int {
	if status, ok := o[common.ResponseObjectFieldStatusCode]; ok {
		return int(status.(float64))
	}
	panic("stretchr: Failed to get status code from the response object, and all responses should have a status code.")
}

// Context returns the context in the response if present, otherwise it returns an empty string.
func (o ResponseObject) Context() string {
	if context, ok := o[common.ResponseObjectFieldContext].(string); ok {
		return context
	}
	return ""
}

// Data returns the ~data object from the response
func (o ResponseObject) Data() interface{} {
	return o[common.ResponseObjectFieldData]
}

// Errors returns the array of errors returned by Stretchr or an empty array if no errors are present.
func (o ResponseObject) Errors() []string {

	if _, ok := o[responseObjectFieldErrorStrings]; !ok {

		errorData, hasErrors := o[common.ResponseObjectFieldErrors]

		// if no errors, return early
		if !hasErrors || errorData == nil {
			o[responseObjectFieldErrorStrings] = []string{}
			return o[responseObjectFieldErrorStrings].([]string)
		}

		errorDataArray := errorData.([]interface{})

		errorStrings := make([]string, len(errorDataArray))
		for i, e := range errorDataArray {
			errorStrings[i] = e.(map[string]interface{})[common.ResponseObjectFieldErrorsMessage].(string)
		}

		o[responseObjectFieldErrorStrings] = errorStrings

	}

	return o[responseObjectFieldErrorStrings].([]string)

}

// HasErrors determines if the response contains errors sent by stretchr
func (o ResponseObject) HasErrors() bool {
	if _, hasErrors := o[common.ResponseObjectFieldErrors]; hasErrors {
		return true
	}
	return false
}

// ChangeInfo returns the ChangeInfo object from the response
func (o ResponseObject) ChangeInfo() ChangeInfo {
	if changeInfo, hasChangeInfo := o[common.ResponseObjectFieldChangeInfo]; hasChangeInfo {
		return ChangeInfo(changeInfo.(map[string]interface{}))
	}
	return NoChangeInfo
}
