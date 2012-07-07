package stretchr

import (
	"testing"
)

func getTestAction(method, path string) *Action {
	s := InProject("test")
	a := s.MakeAction(method, path)
	return a
}

func TestMakeAction(t *testing.T) {

	s := InProject("test")
	q := s.MakeAction(Get, "people")

	AssertEqual(t, Get, q.Method)
	AssertEqual(t, "people", q.Path)
	AssertEqual(t, s, q.Session)

}

func TestAction_GetRequest_Read(t *testing.T) {

	a := getTestAction(Get, "people/123")
	r, _ := a.GetRequest()

	AssertRequest(t, r, Get, "people/123")
	AssertNoRequestBody(t, r)

}

func TestAction_GetRequest_ReadMany(t *testing.T) {

	a := getTestAction(Get, "people")
	r, _ := a.GetRequest()

	AssertRequest(t, r, Get, "people")
	AssertNoRequestBody(t, r)

}

func TestAction_GetRequest_Create(t *testing.T) {

	data := "{\"name\":\"Mat\"}"
	a := getTestAction(Post, "people").WithData(data)
	r, _ := a.GetRequest()

	AssertRequest(t, r, Post, "people")
	AssertRequestBody(t, r, data)

}

func TestAction_GetRequest_Update(t *testing.T) {

	data := "{\"name\":\"Mat\"}"
	a := getTestAction(Put, "people/123").WithData(data)
	r, _ := a.GetRequest()

	AssertRequest(t, r, Put, "people/123")
	AssertRequestBody(t, r, data)

}

func TestAction_GetRequest_Replace(t *testing.T) {

	data := "{\"name\":\"Mat\"}"
	a := getTestAction(Post, "people/123").WithData(data)
	r, _ := a.GetRequest()

	AssertRequest(t, r, Post, "people/123")
	AssertRequestBody(t, r, data)

}

func TestAction_GetRequest_Delete(t *testing.T) {

	a := getTestAction(Delete, "people/123")
	r, _ := a.GetRequest()

	AssertRequest(t, r, Delete, "people/123")
	AssertNoRequestBody(t, r)

}

func TestAction_GetRequest_DeleteMany(t *testing.T) {

	a := getTestAction(Delete, "people")
	r, _ := a.GetRequest()

	AssertRequest(t, r, Delete, "people")
	AssertNoRequestBody(t, r)

}
