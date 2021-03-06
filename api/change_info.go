package api

import (
	"github.com/stretchr/sdk-go/common"
)

// ChangeInfo represents a data object containing information about
// what changed in response to a request.
type ChangeInfo map[string]interface{}

// NoChangeInfo is a nil ChangeInfo object, and a useful shortcut.
var NoChangeInfo ChangeInfo

// Created gets the number of records that were created as indicated by
// this ChangeInfo object.
func (c ChangeInfo) Created() int {
	if val, ok := c[common.ChangeInfoPublicFieldCreated]; ok {
		return int(val.(float64))
	}
	return 0
}

// Updated gets the number of records that were updated as indicated by
// this ChangeInfo object.
func (c ChangeInfo) Updated() int {
	if val, ok := c[common.ChangeInfoPublicFieldUpdated]; ok {
		return int(val.(float64))
	}
	return 0
}

// Deleted gets the number of records that were deleted as indicated by
// this ChangeInfo object.
func (c ChangeInfo) Deleted() int {
	if val, ok := c[common.ChangeInfoPublicFieldDeleted]; ok {
		return int(val.(float64))
	}
	return 0
}

// HasDeltas gets whether deltas exist in this ChangeInfo object
func (c ChangeInfo) HasDeltas() bool {
	if val, ok := c[common.ChangeInfoPublicFieldDeltas]; ok {
		if len(val.([]interface{})) == 0 {
			return false
		}
		return true
	}
	return false
}

// Deltas gets the array of (map[string]interface{}) Deltas that were created in the last
// request if any.
func (c ChangeInfo) Deltas() []map[string]interface{} {
	deltas := []map[string]interface{}{}
	if val, ok := c[common.ChangeInfoPublicFieldDeltas]; ok {
		for _, delta := range val.([]interface{}) {
			deltas = append(deltas, delta.(map[string]interface{}))
		}
	}
	return deltas
}
