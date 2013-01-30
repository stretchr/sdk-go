package api

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestResponseObject_StatusCode(t *testing.T) {

	obj := ResponseObject(map[string]interface{}{"~s": float64(200)})

	assert.Equal(t, 200, obj.StatusCode(), "StatusCode")

}

func TestResponseObject_Context(t *testing.T) {

	obj := ResponseObject(map[string]interface{}{"~x": "context"})

	assert.Equal(t, "context", obj.Context(), "Context")

}

func TestResponseObject_Context_NoContext(t *testing.T) {

	obj := ResponseObject(map[string]interface{}{"nope": "context"})

	assert.Equal(t, "", obj.Context(), "Context")

}

func TestResponseObject_Data(t *testing.T) {

	obj := ResponseObject(map[string]interface{}{"~d": map[string]interface{}{"name": "Mat"}})

	assert.Equal(t, "Mat", obj.Data().(map[string]interface{})["name"])

}

func TestResponseObject_Data_WithNoData(t *testing.T) {

	obj := ResponseObject(map[string]interface{}{"nope": map[string]interface{}{"name": "Mat"}})

	assert.Nil(t, obj.Data())

}

func TestResponseObject_Errors(t *testing.T) {

	obj := ResponseObject(map[string]interface{}{"~e": []interface{}{map[string]interface{}{"~m": "Something went wrong"}}})

	errs := obj.Errors()
	if assert.Equal(t, 1, len(errs)) {
		assert.Equal(t, "Something went wrong", errs[0])
	}

}

func TestResponseObject_ChangeInfo(t *testing.T) {

	obj := ResponseObject(map[string]interface{}{"~ch": map[string]interface{}{"~c": float64(1), "~u": float64(2), "~d": float64(3), "~ids": []interface{}{"ABC", "DEF"}}})

	changeInfo := obj.ChangeInfo()

	if assert.NotNil(t, changeInfo) {

		assert.Equal(t, 1, changeInfo.Created())
		assert.Equal(t, 2, changeInfo.Updated())
		assert.Equal(t, 3, changeInfo.Deleted())

		ids := changeInfo.IDs()

		if assert.Equal(t, 2, len(ids)) {
			assert.Equal(t, "ABC", ids[0])
			assert.Equal(t, "DEF", ids[1])
		}

	}

}

func TestResponseObject_ChangeInfo_NoChangeInfo(t *testing.T) {

	obj := ResponseObject(map[string]interface{}{"nope": map[string]interface{}{"~c": float64(1), "~u": float64(2), "~d": float64(3), "~ids": []interface{}{"ABC", "DEF"}}})

	changeInfo := obj.ChangeInfo()

	assert.Equal(t, NoChangeInfo, changeInfo)

}
