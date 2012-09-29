package stretchr

import (
	"testing"
)

func TestSession_Many(t *testing.T) {

	m := TestSession.Many("people")

	if m == nil {
		t.Error("TestSession.Many shouldn't return nil")
	} else {

		AssertEqual(t, TestSession, m.session)
		AssertEqual(t, "people", m.path)

	}

}

func TestSession_Manyf(t *testing.T) {

	m := TestSession.Manyf("people/%s/books", "123")

	if m == nil {
		t.Error("TestSession.Manyf shouldn't return nil")
	} else {

		AssertEqual(t, TestSession, m.session)
		AssertEqual(t, "people/123/books", m.path)

	}

}
