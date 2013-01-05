package stretchr

import (
	"errors"
	"github.com/stretchrcom/stretchr-sdk-go/api"
)

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
