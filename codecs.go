package stretchr

import (
	"encoding/json"
)

/*
	TODO: replace this after opening stretchrcom/codecs package
	Search code tag: #codecs
*/

func ObjectToBytes(object interface{}) ([]byte, error) {
	return json.Marshal(object)
}
func BytesToObject(bytes []byte) (interface{}, error) {
	var object interface{}
	err := json.Unmarshal(bytes, object)
	return object, err
}
