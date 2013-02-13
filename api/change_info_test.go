package api

import (
	"github.com/stretchrcom/sdk-go/common"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestChangeInfo_Changed(t *testing.T) {

	changeInfo := ChangeInfo(map[string]interface{}{ChangeInfoPublicFieldCreated: float64(2)})

	assert.Equal(t, 2, changeInfo.Created())

}

func TestChangeInfo_Updated(t *testing.T) {

	changeInfo := ChangeInfo(map[string]interface{}{ChangeInfoPublicFieldUpdated: float64(25)})

	assert.Equal(t, 25, changeInfo.Updated())

}

func TestChangeInfo_Deleted(t *testing.T) {

	changeInfo := ChangeInfo(map[string]interface{}{ChangeInfoPublicFieldDeleted: float64(26)})

	assert.Equal(t, 26, changeInfo.Deleted())

}

func TestChangeInfo_Deltas(t *testing.T) {

	changeInfo := ChangeInfo(map[string]interface{}{ChangeInfoPublicFieldDeltas: []interface{}{map[string]interface{}{common.DataFieldID: "one"}, map[string]interface{}{common.DataFieldID: "two"}, map[string]interface{}{common.DataFieldID: "three"}}})

	assert.Equal(t, "one", changeInfo.Deltas()[0][common.DataFieldID])
	assert.Equal(t, "two", changeInfo.Deltas()[1][common.DataFieldID])
	assert.Equal(t, "three", changeInfo.Deltas()[2][common.DataFieldID])

}

func TestChangeInfo_GettersWithNoData(t *testing.T) {

	changeInfo := ChangeInfo{}

	assert.Equal(t, 0, changeInfo.Created())
	assert.Equal(t, 0, changeInfo.Updated())
	assert.Equal(t, 0, changeInfo.Deleted())
	assert.Equal(t, 0, len(changeInfo.Deltas()))

}
