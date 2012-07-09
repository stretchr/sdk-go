package stretchr

import (
	"testing"
)

func TestMakeResource(t *testing.T) {

	r := MakeResource(TestSession, "people")

	if r == nil {
		t.Error("MakeResource shouldn't return nil.")
	} else {

		AssertEqual(t, "people", r.path)
		AssertEqual(t, TestSession, r.session)

	}

}

func TestPath(t *testing.T) {

	r := MakeResource(TestSession, "people")
	AssertEqual(t, "people", r.Path())

	r = MakeResource(TestSession, "people/ABC/books/DEF")
	AssertEqual(t, "people/ABC/books/DEF", r.Path())

}

func TestGetAndSet(t *testing.T) {

	r := MakeResource(TestSession, "people")

	/*
		Set
	*/
	AssertEqual(t, r, r.Set("name", "Mat"))
	AssertEqual(t, "Mat", r.data["name"])

	/*
		Get
	*/
	r.data["name"] = "Laurie"
	AssertEqual(t, "Laurie", r.Get("name"))

}

func TestRemove(t *testing.T) {

	r := MakeResource(TestSession, "people")

	r.Set("name", "Mat")

	AssertEqual(t, "Mat", r.data["name"])

	r.Remove("name")

	AssertEqual(t, nil, r.data["name"])

}

func TestGetAndSetID(t *testing.T) {

	r := MakeResource(TestSession, "people")

	/*
		Get with no ID
	*/
	AssertEqual(t, EmptyID, r.GetID())

	/*
		Set
	*/
	AssertEqual(t, r, r.SetID("ABC"))
	AssertEqual(t, "ABC", r.data[IDKey])

	/*
		Get
	*/
	r.data["~id"] = "DEF"
	AssertEqual(t, "DEF", r.GetID())
	AssertEqual(t, "DEF", r.Get(IDKey))

	/*
		Clear ID
	*/
	AssertEqual(t, r, r.ClearID())
	AssertEqual(t, EmptyID, r.GetID())
	AssertEqual(t, nil, r.data[IDKey])

}

func TestData(t *testing.T) {

	r := MakeResource(TestSession, "people")
	r.Set("name", "Mat").Set("age", 29)

	var data map[string]interface{} = r.Data()

	AssertEqual(t, "Mat", data["name"])
	AssertEqual(t, 29, data["age"])

}

func TestAbsoluteURL(t *testing.T) {

	r := MakeResource(TestSession, "people/ABC/books/DEF")
	AssertEqual(t, "http://test.stretchr.com/api/v1/people/ABC/books/DEF", r.AbsoluteURL())

}
