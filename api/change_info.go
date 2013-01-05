package api

const (
	ChangeInfoFieldCreated string = "~c"
	ChangeInfoFieldUpdated string = "~u"
	ChangeInfoFieldDeleted string = "~d"
	ChangeInfoFieldIDs     string = "~ids"
)

type ChangeInfo map[string]interface{}

var NoChangeInfo ChangeInfo = nil

func (c ChangeInfo) Created() int {
	if val, ok := c[ChangeInfoFieldCreated]; ok {
		return int(val.(float64))
	}
	return 0
}

func (c ChangeInfo) Updated() int {
	if val, ok := c[ChangeInfoFieldUpdated]; ok {
		return int(val.(float64))
	}
	return 0
}

func (c ChangeInfo) Deleted() int {
	if val, ok := c[ChangeInfoFieldDeleted]; ok {
		return int(val.(float64))
	}
	return 0
}

func (c ChangeInfo) IDs() []string {
	ids := []string{}
	if val, ok := c[ChangeInfoFieldIDs]; ok {
		for _, id := range val.([]interface{}) {
			ids = append(ids, id.(string))
		}
	}
	return ids
}
