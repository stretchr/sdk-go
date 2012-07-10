package stretchr

import (
	"net/http"
	"testing"
)

func TestMakeMany(t *testing.T) {

	m := makeMany(TestSession, "people")

	if m == nil {
		t.Error("MakeMany shouldn't return nil")
	} else {

		AssertEqual(t, TestSession, m.session)
		AssertEqual(t, "people", m.path)

	}

}

func TestManyPath(t *testing.T) {

	m := makeMany(TestSession, "people")
	AssertEqual(t, "people", m.Path())

}

func TestMany_SetParameter(t *testing.T) {

	m := makeMany(TestSession, "people")
	AssertEqual(t, m, m.SetParameter("~monkey", "10"))
	AssertEqual(t, "people?~monkey=10", m.Path())

}

func TestMany_RemoveParameter(t *testing.T) {

	m := makeMany(TestSession, "people").SetParameter("~monkey", "100")
	AssertEqual(t, m, m.SetParameter("~limit", "10"))

	AssertContains(t, m.Path(), "people?")
	AssertContains(t, m.Path(), "~monkey=100")
	AssertContains(t, m.Path(), "~limit=10")

	AssertEqual(t, m, m.RemoveParameter("~monkey"))
	AssertEqual(t, "people?~limit=10", m.Path())

}

func TestMany_Parameters(t *testing.T) {

	m := makeMany(TestSession, "people").SetParameter("~monkey", "100")

	AssertEqual(t, m.parameters.Get("~monkey"), m.Parameters().Get("~monkey"))

}

func TestMany_Limit(t *testing.T) {

	m := makeMany(TestSession, "people")
	AssertEqual(t, m, m.Limit(10))
	AssertEqual(t, "people?~limit=10", m.Path())

}

func TestMany_Skip(t *testing.T) {

	m := makeMany(TestSession, "people")
	AssertEqual(t, m, m.Skip(10))
	AssertEqual(t, "people?~skip=10", m.Path())

}

func TestMany_Page(t *testing.T) {

	m := makeMany(TestSession, "people")
	AssertEqual(t, m, m.Page(2, 10))
	AssertContains(t, m.Path(), "people?")
	AssertContains(t, m.Path(), "~skip=10")
	AssertContains(t, m.Path(), "~limit=10")

	AssertEqual(t, m, m.Page(3, 10))
	AssertContains(t, m.Path(), "people?")
	AssertContains(t, m.Path(), "~skip=20")
	AssertContains(t, m.Path(), "~limit=10")

	AssertEqual(t, m, m.Page(1, 5))
	AssertEqual(t, "people?~limit=5", m.Path())

}

func TestManyRead(t *testing.T) {

	// use the test requester
	ActiveRequester = ActiveTestRequester.Reset()

	responseData := []map[string]interface{}{map[string]interface{}{"name": "Mat", "age": 29, "developer": true, IDKey: "ABC"}, map[string]interface{}{"name": "Laurie", "age": 28, "developer": false, IDKey: "DEF"}}
	ActiveTestRequester.ResponseToReturn = MakeTestResponseWithData(http.StatusOK, makeStandardResponseObject(http.StatusOK, responseData))

	// make a resource
	resourceCollection, err := makeMany(TestSession, "people").Read()

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
