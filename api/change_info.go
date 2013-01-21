package api

const (
	ChangeInfoFieldCreated string = "~c"
	ChangeInfoFieldUpdated string = "~u"
	ChangeInfoFieldDeleted string = "~d"
	ChangeInfoFieldIDs     string = "~ids"
)

// ChangeInfo represents a data object containing information about
// what changed in response to a request.
type ChangeInfo map[string]interface{}

// NoChangeInfo is a nil ChangeInfo object, and a useful shortcut.
var NoChangeInfo ChangeInfo = nil

// Created gets the number of records that were created as indicated by
// this ChangeInfo object.
func (c ChangeInfo) Created() int {
	if val, ok := c[ChangeInfoFieldCreated]; ok {
		return int(val.(float64))
	}
	return 0
}

// Updated gets the number of records that were updated as indicated by
// this ChangeInfo object.
func (c ChangeInfo) Updated() int {
	if val, ok := c[ChangeInfoFieldUpdated]; ok {
		return int(val.(float64))
	}
	return 0
}

// Deleted gets the number of records that were deleted as indicated by
// this ChangeInfo object.
func (c ChangeInfo) Deleted() int {
	if val, ok := c[ChangeInfoFieldDeleted]; ok {
		return int(val.(float64))
	}
	return 0
}

// IDs gets the array of (string) IDs that were created in the last
// request if any.
func (c ChangeInfo) IDs() []string {
	ids := []string{}
	if val, ok := c[ChangeInfoFieldIDs]; ok {
		for _, id := range val.([]interface{}) {
			ids = append(ids, id.(string))
		}
	}
	return ids
}
