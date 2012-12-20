package stretchr

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestResourceObject_Interface(t *testing.T) {

	assert.Implements(t, (*Resource)(nil), new(ResourceObject))

}

func TestMakeResource(t *testing.T) {

	r := MakeResourceAt("people/123")
	r.Data["name"] = "Mat"

	if assert.NotNil(t, r) {
		assert.Equal(t, "people/123", r.ResourcePath())
		assert.Equal(t, "Mat", r.ResourceData()["name"])
	}

}

func TestResourceObject_ResourcePath_GetsIDFromBody(t *testing.T) {

	r := MakeResourceAt("people")
	r.Data[dataFieldID] = "123"

	assert.Equal(t, "people/123", r.ResourcePath())

}
func TestResourceObject_ResourcePath_ReplacesIDFromBody(t *testing.T) {

	r := MakeResourceAt("people/abc")
	r.Data[dataFieldID] = "123"

	assert.Equal(t, "people/123", r.ResourcePath())

}