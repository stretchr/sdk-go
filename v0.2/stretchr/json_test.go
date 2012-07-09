package stretchr

import (
	"testing"
)

func TestToJson(t *testing.T) {

	data := make(map[string]interface{})
	data["name"] = "Mat"
	data["age"] = 29
	data["developer"] = true

	json, err := ToJson(data)

	if err != nil {
		t.Errorf("ToJson shouldn't throw error: %s", err)
	} else {

		AssertEqual(t, "{\"age\":29,\"developer\":true,\"name\":\"Mat\"}", json)

	}

}

func TestFromJson(t *testing.T) {

	json := "{\"age\":29,\"developer\":true,\"name\":\"Mat\"}"

	obj, err := FromJson(json)

	if err != nil {
		t.Errorf("FromJson shouldn't throw error: %s", err)
	} else {

		AssertEqual(t, "Mat", obj["name"])
		AssertEqual(t, float64(29), obj["age"])
		AssertEqual(t, true, obj["developer"])

	}

}
