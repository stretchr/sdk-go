package stretchr

import (
	"testing"
)

func TestRequestForRead(t *testing.T) {
	
	r, _ := RequestForRead(testSession(), "people", "123")
	
	AssertRequest(t, r, Get, "people/123")
	AssertNoRequestBody(t, r)
	
}

func TestRequestForReadMany(t *testing.T) {
	
	r, _ := RequestForReadMany(testSession(), "people")
	
	AssertRequest(t, r, Get, "people")
	AssertNoRequestBody(t, r)
	
}

func TestRequestForCreate(t *testing.T) {
	
	data := "{\"name\":\"Mat\"}"
	r, _ := RequestForCreate(testSession(), "people", data)
	
	AssertRequest(t, r, Post, "people")
	AssertRequestBody(t, r, data)
	
}

func TestRequestForUpdate(t *testing.T) {
	
	data := "{\"name\":\"Mat\"}"
	r, _ := RequestForUpdate(testSession(), "people", "ABC", data)
	
	AssertRequest(t, r, Put, "people/ABC")
	AssertRequestBody(t, r, data)
	
}

func TestRequestForReplace(t *testing.T) {
	
	data := "{\"name\":\"Mat\"}"
	r, _ := RequestForReplace(testSession(), "people", "ABC", data)
	
	AssertRequest(t, r, Post, "people/ABC")
	AssertRequestBody(t, r, data)
	
}

func TestRequestForDelete(t *testing.T) {
	
	r, _ := RequestForDelete(testSession(), "people", "123")
	
	AssertRequest(t, r, Delete, "people/123")
	AssertNoRequestBody(t, r)
	
}

func TestRequestForDeleteMany(t *testing.T) {
	
	r, _ := RequestForDeleteMany(testSession(), "people")
	
	AssertRequest(t, r, Delete, "people")
	AssertNoRequestBody(t, r)
	
}
