package stretchr

import (
	"testing"
)

func TesttoJson(t *testing.T) {

	data := make(map[string]interface{})
	data["name"] = "Mat"
	data["age"] = 29
	data["developer"] = true

	json, err := toJson(data)

	if err != nil {
		t.Errorf("toJson shouldn't throw error: %s", err)
	} else {

		AssertEqual(t, "{\"age\":29,\"developer\":true,\"name\":\"Mat\"}", json)

	}

}

func TestfromJson(t *testing.T) {

	json := "{\"age\":29,\"developer\":true,\"name\":\"Mat\"}"

	obj, err := fromJson(json)

	if err != nil {
		t.Errorf("fromJson shouldn't throw error: %s", err)
	} else {

		AssertEqual(t, "Mat", obj["name"])
		AssertEqual(t, float64(29), obj["age"])
		AssertEqual(t, true, obj["developer"])

	}

}
