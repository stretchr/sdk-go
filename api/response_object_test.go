package api

import (
	"github.com/stretchrcom/sdk-go/common"
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

	obj := ResponseObject(map[string]interface{}{ResponseObjectFieldData: map[string]interface{}{"name": "Mat"}})

	assert.Equal(t, "Mat", obj.Data().(map[string]interface{})["name"])

}

func TestResponseObject_Data_WithNoData(t *testing.T) {

	obj := ResponseObject(map[string]interface{}{"nope": map[string]interface{}{"name": "Mat"}})

	assert.Nil(t, obj.Data())

}

func TestResponseObject_Errors(t *testing.T) {

	obj := ResponseObject(map[string]interface{}{ResponseObjectFieldErrors: []interface{}{map[string]interface{}{"~m": "Something went wrong"}}})

	errs := obj.Errors()
	if assert.Equal(t, 1, len(errs)) {
		assert.Equal(t, "Something went wrong", errs[0])
	}

}

func TestResponseObject_ChangeInfo(t *testing.T) {

	obj := ResponseObject(map[string]interface{}{ResponseObjectFieldChangeInfo: map[string]interface{}{ChangeInfoFieldCreated: float64(1), ChangeInfoFieldUpdated: float64(2), ChangeInfoFieldDeleted: float64(3), ChangeInfoFieldDeltas: []interface{}{map[string]interface{}{common.DataFieldID: "ABC"}, map[string]interface{}{common.DataFieldID: "DEF"}}}})

	changeInfo := obj.ChangeInfo()

	if assert.NotNil(t, changeInfo) {

		assert.Equal(t, 1, changeInfo.Created())
		assert.Equal(t, 2, changeInfo.Updated())
		assert.Equal(t, 3, changeInfo.Deleted())

		deltas := changeInfo.Deltas()

		if assert.Equal(t, 2, len(deltas)) {
			assert.Equal(t, "ABC", deltas[0][common.DataFieldID])
			assert.Equal(t, "DEF", deltas[1][common.DataFieldID])
		}

	}

}

func TestResponseObject_ChangeInfo_NoChangeInfo(t *testing.T) {

	obj := ResponseObject(map[string]interface{}{"nope": map[string]interface{}{ChangeInfoFieldCreated: float64(1), ChangeInfoFieldUpdated: float64(2), ChangeInfoFieldDeleted: float64(3), "~ids": []interface{}{"ABC", "DEF"}}})

	changeInfo := obj.ChangeInfo()

	assert.Equal(t, NoChangeInfo, changeInfo)

}
