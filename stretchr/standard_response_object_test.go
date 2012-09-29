package stretchr

import (
	"fmt"
	"net/http"
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

func TestExtractStandardResponseObject_ResponseBody(t *testing.T) {

	data := map[string]interface{}{"name": "Mat", "age": 29}
	response := MakeTestResponseWithData(200, makeStandardResponseObject(200, data))

	obj, err := ExtractStandardResponseObject(response)

	if err != nil {
		t.Errorf("ExtractStandardResponseObject shouldn't raise error: %s", err)
	}

	AssertEqual(t, "{\"d\":{\"age\":29,\"name\":\"Mat\"},\"s\":200,\"w\":true}", obj.ResponseBody)

}

func TestExtractStandardResponseObject_SingleDataObject(t *testing.T) {

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

func TestExtractStandardResponseObject_DataCollection(t *testing.T) {

	data1 := map[string]interface{}{"name": "Mat", "age": 29}
	data2 := map[string]interface{}{"name": "Chris", "age": 29}
	data3 := map[string]interface{}{"name": "Laurie", "age": 29}
	data := map[string]interface{}{"c": 3, "i": []interface{}{data1, data2, data3}}

	response := MakeTestResponseWithData(200, makeStandardResponseObject(200, data))

	obj, err := ExtractStandardResponseObject(response)

	if err != nil {
		t.Errorf("ExtractStandardResponseObject shouldn't raise error: %s", err)
	}

	AssertEqual(t, 200, obj.StatusCode)
	AssertEqual(t, true, obj.Worked)
	AssertEqual(t, 0, len(obj.Errors))
	AssertEqual(t, float64(3), obj.Data["c"])

	AssertEqual(t, "Mat", obj.Data["i"].([]interface{})[0].(map[string]interface{})["name"])
	AssertEqual(t, "Chris", obj.Data["i"].([]interface{})[1].(map[string]interface{})["name"])
	AssertEqual(t, "Laurie", obj.Data["i"].([]interface{})[2].(map[string]interface{})["name"])

}

func TestStandardResponseObject_Errors(t *testing.T) {

	var message string = "Test Error Message"
	response := MakeTestResponseWithData(500, makeFailedStandardResponseObject(500, []string{message}))

	obj, err := ExtractStandardResponseObject(response)

	if err != nil {
		t.Errorf("ExtractStandardResponseObject shouldn't raise error: %s", err)
	}

	AssertEqual(t, 500, obj.StatusCode)
	AssertEqual(t, false, obj.Worked)
	if AssertEqual(t, 1, len(obj.Errors)) {
		AssertEqual(t, message, obj.Errors[0].(map[string]interface{})["Message"])
	}

}

func TestStandardResponseObject_GetError(t *testing.T) {

	sro := &StandardResponseObject{500, []interface{}{map[string]interface{}{"Message": "Something went wrong :-("}}, false, nil, "", ""}

	AssertEqual(t, "Something went wrong :-(", fmt.Sprintf("%s", sro.GetError()))

	sro = &StandardResponseObject{500, make([]interface{}, 0), false, nil, "", ""}

	AssertEqual(t, UnknownError, sro.GetError())

}

func TestStandardResponseObject_GetError_NotFound(t *testing.T) {

	sro := &StandardResponseObject{http.StatusNotFound, []interface{}{map[string]interface{}{"Message": "Something went wrong :-("}}, false, nil, "", ""}

	AssertEqual(t, NotFound, sro.GetError())

}
