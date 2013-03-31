package api

const (
	ResponseObjectFieldData          string = "~d"
	ResponseObjectFieldChangeInfo    string = "~ch"
	ResponseObjectFieldErrors        string = "~e"
	ResponseObjectFieldStatusCode    string = "~s"
	ResponseObjectFieldContext       string = "~x"
	ResponseObjectFieldErrorsMessage string = "~m"
	responseObjectFieldErrorStrings  string = "_e"
)

type ResponseObject map[string]interface{}

func (o ResponseObject) StatusCode() int {
	if status, ok := o[ResponseObjectFieldStatusCode]; ok {
		return int(status.(float64))
	}
	panic("stretchr: Failed to get status code from the response object, and all responses should have a status code.")
}

func (o ResponseObject) Context() string {
	if context, ok := o[ResponseObjectFieldContext].(string); ok {
		return context
	}
	return ""
}

func (o ResponseObject) Data() interface{} {
	return o[ResponseObjectFieldData]
}

func (o ResponseObject) Errors() []string {

	if _, ok := o[responseObjectFieldErrorStrings]; !ok {

		errorData, hasErrors := o[ResponseObjectFieldErrors]

		// if no errors, return early
		if !hasErrors || errorData == nil {
			o[responseObjectFieldErrorStrings] = []string{}
			return o[responseObjectFieldErrorStrings].([]string)
		}

		errorDataArray := errorData.([]interface{})

		errorStrings := make([]string, len(errorDataArray))
		for i, e := range errorDataArray {
			errorStrings[i] = e.(map[string]interface{})[ResponseObjectFieldErrorsMessage].(string)
		}

		o[responseObjectFieldErrorStrings] = errorStrings

	}

	return o[responseObjectFieldErrorStrings].([]string)

}

func (o ResponseObject) HasErrors() bool {
	if _, hasErrors := o[ResponseObjectFieldErrors]; hasErrors {
		return true
	}
	return false
}

func (o ResponseObject) ChangeInfo() ChangeInfo {
	if changeInfo, hasChangeInfo := o[ResponseObjectFieldChangeInfo]; hasChangeInfo {
		return ChangeInfo(changeInfo.(map[string]interface{}))
	}
	return NoChangeInfo
}
