package stretchr

import (
	"errors"
	"github.com/stretchrcom/stretchr-sdk-go/api"
)

var (
	ErrSingleObjectExpectedButGotArray         = errors.New("stretchr: Array in response data when a single object was expected.")
	ErrSingleObjectExpectedButGotNil           = errors.New("stretchr: Nil in response data when a single object was expected.")
	ErrSingleObjectExpectedButGotSomethingElse = errors.New("stretchr: Unexpected thing in response data when a single object was expected.")
)

// GetErrorsFromResponseObject gets the array of errors from the ResponseObject.
func GetErrorsFromResponseObject(response api.ResponseObject) []error {

	errorStrings := response.Errors()
	if len(errorStrings) > 0 {

		errorArray := make([]error, len(errorStrings))
		for errIndex, errString := range errorStrings {
			errorArray[errIndex] = errors.New(errString)
		}

		return errorArray

	}

	return []error{}

}

// GetErrorFromResponseObject gets the first error from the ResponseObject, or returns nil
// if there were no errors.
func GetErrorFromResponseObject(response api.ResponseObject) error {
	errs := GetErrorsFromResponseObject(response)
	if len(errs) == 0 {
		return nil
	}
	return errs[0]
}
