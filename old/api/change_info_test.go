package api

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestChangeInfo_Changed(t *testing.T) {

	changeInfo := ChangeInfo(map[string]interface{}{"~c": float64(2)})

	assert.Equal(t, 2, changeInfo.Created())

}

func TestChangeInfo_Updated(t *testing.T) {

	changeInfo := ChangeInfo(map[string]interface{}{"~u": float64(25)})

	assert.Equal(t, 25, changeInfo.Updated())

}

func TestChangeInfo_Deleted(t *testing.T) {

	changeInfo := ChangeInfo(map[string]interface{}{"~d": float64(26)})

	assert.Equal(t, 26, changeInfo.Deleted())

}

func TestChangeInfo_IDs(t *testing.T) {

	changeInfo := ChangeInfo(map[string]interface{}{"~ids": []interface{}{"one", "two", "three"}})

	assert.Equal(t, "one", changeInfo.IDs()[0])
	assert.Equal(t, "two", changeInfo.IDs()[1])
	assert.Equal(t, "three", changeInfo.IDs()[2])

}

func TestChangeInfo_GettersWithNoData(t *testing.T) {

	changeInfo := ChangeInfo{}

	assert.Equal(t, 0, changeInfo.Created())
	assert.Equal(t, 0, changeInfo.Updated())
	assert.Equal(t, 0, changeInfo.Deleted())
	assert.Equal(t, 0, len(changeInfo.IDs()))

}
