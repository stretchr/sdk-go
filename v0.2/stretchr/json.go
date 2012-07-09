package stretchr

import (
	"encoding/json"
)

// toJson turns a map[string]interface{} object into a JSON string, or
// returns an error if something goes wrong.
func toJson(data map[string]interface{}) (string, error) {

	bytes, err := json.Marshal(data)

	if err != nil {
		return "", err
	}

	return string(bytes), nil

}

// fromJson turns the JSON string of an object (not array) into an object,
// or returns an error if something goes wrong.
func fromJson(j string) (map[string]interface{}, error) {
	var obj map[string]interface{}
	err := json.Unmarshal([]byte(j), &obj)
	return obj, err
}
