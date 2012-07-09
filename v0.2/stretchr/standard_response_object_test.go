package stretchr

import (
	"testing"
)

func makeStandardResponseObject_internal(statusCode int, data interface{}, errors []string) map[string]interface{} {

	obj := make(map[string]interface{})

	if data != nil {
		obj["d"] = data
	}
	obj["s"] = statusCode
	obj["w"] = statusCode >= 200 && statusCode <= 299

	if len(errors) > 0 {
		errlist := make([]interface{}, len(errors))
		for errIndex, err := range errors {
			errlist[errIndex] = map[string]interface{}{"Message": err}
		}
		obj["e"] = errlist
	}

	return obj

}

func makeStandardResponseObject(statusCode int, data interface{}) map[string]interface{} {
	return makeStandardResponseObject_internal(statusCode, data, []string{})
}
func makeFailedStandardResponseObject(statusCode int, errors []string) map[string]interface{} {
	return makeStandardResponseObject_internal(statusCode, nil, errors)
}

func TestExtractStandardResponseObject(t *testing.T) {

	data := map[string]interface{}{"name": "Mat", "age": 29}
	response := MakeTestResponseWithData(200, makeStandardResponseObject(200, data))

	obj, err := ExtractStandardResponseObject(response)

	if err != nil {
		t.Errorf("ExtractStandardResponseObject shouldn't raise error: %s", err)
	}

	AssertEqual(t, 200, obj.StatusCode)
	AssertEqual(t, true, obj.Worked)
	AssertEqual(t, 0, len(obj.Errors))
	AssertEqual(t, "Mat", obj.Data["name"])
	AssertEqual(t, float64(29), obj.Data["age"])

}
