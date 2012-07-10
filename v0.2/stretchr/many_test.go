package stretchr

import (
	"net/http"
	"testing"
)

func TestMakeMany(t *testing.T) {

	m := MakeMany(TestSession, "people")

	if m == nil {
		t.Error("MakeMany shouldn't return nil")
	} else {

		AssertEqual(t, TestSession, m.session)
		AssertEqual(t, "people", m.path)

	}

}

func TestManyRead(t *testing.T) {

	// use the test requester
	ActiveRequester = ActiveTestRequester.Reset()

	responseData := []map[string]interface{}{map[string]interface{}{"name": "Mat", "age": 29, "developer": true, IDKey: "ABC"}, map[string]interface{}{"name": "Laurie", "age": 28, "developer": false, IDKey: "DEF"}}
	ActiveTestRequester.ResponseToReturn = MakeTestResponseWithData(http.StatusOK, makeStandardResponseObject(http.StatusOK, responseData))

	// make a resource
	resourceCollection, err := MakeMany(TestSession, "people").Read()

	if err != nil {
		t.Errorf("Shouldn't throw error: %s", err)
	}

	AssertLastRequest(t, ReadMethod, TestSession.Url("people"), "", "PRIVATE")

	if resourceCollection == nil {
		t.Errorf("MakeMany().Read() shouldn't be nil")
	} else {

		if AssertEqual(t, 2, len(resourceCollection.Resources)) {

			AssertEqual(t, "Mat", resourceCollection.Resources[0].Get("name"))
			AssertEqual(t, float64(29), resourceCollection.Resources[0].Get("age"))
			AssertEqual(t, true, resourceCollection.Resources[0].Get("developer"))
			AssertEqual(t, "ABC", resourceCollection.Resources[0].GetID())

			AssertEqual(t, "Laurie", resourceCollection.Resources[1].Get("name"))
			AssertEqual(t, float64(28), resourceCollection.Resources[1].Get("age"))
			AssertEqual(t, false, resourceCollection.Resources[1].Get("developer"))
			AssertEqual(t, "DEF", resourceCollection.Resources[1].GetID())

		}

	}

}
