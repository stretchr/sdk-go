package stretchr

import (
	"testing"
)

func TestInProject(t *testing.T) {

	var s *Session = InProject("test")

	AssertEqual(t, "test", s.Project)
	AssertEqual(t, "", s.PublicKey)
	AssertEqual(t, "", s.PrivateKey)

}

func TestInProject_WithKeys(t *testing.T) {

	var s *Session = InProject("test").WithKeys("ABC", "DEF")

	AssertEqual(t, "test", s.Project)
	AssertEqual(t, "ABC", s.PublicKey)
	AssertEqual(t, "DEF", s.PrivateKey)

}

func TestWithKeys_InProject(t *testing.T) {

	var s *Session = WithKeys("ABC", "DEF").InProject("test")

	AssertEqual(t, "test", s.Project)
	AssertEqual(t, "ABC", s.PublicKey)
	AssertEqual(t, "DEF", s.PrivateKey)

}

func TestBaseUrl(t *testing.T) {

	s := InProject("test").WithKeys("ABC", "DEF")

	AssertEqual(t, "http://test.stretchr.com/api/v1/", s.baseUrl())

}

func TestUrl(t *testing.T) {

	s := InProject("test").WithKeys("ABC", "DEF")
	AssertEqual(t, "http://test.stretchr.com/api/v1/people/123", s.Url("people/123"))

}
