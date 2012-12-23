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

	assert.Equal(t, r.Data[common.DataFieldID], "123", "MakeResourceAt should set the ID in the data")

	if assert.NotNil(t, r) {
		assert.Equal(t, "people/123", r.ResourcePath())
		assert.Equal(t, "Mat", r.ResourceData()["name"])
	}

}

func TestResourceObject_ResourcePath_GetsIDFromBody(t *testing.T) {

	r := MakeResourceAt("people")
	r.Data[common.DataFieldID] = "123"

	assert.Equal(t, "people/123", r.ResourcePath())

}
func TestResourceObject_ResourcePath_ReplacesIDFromBody(t *testing.T) {

	r := MakeResourceAt("people/abc")
	r.Data[common.DataFieldID] = "123"

	assert.Equal(t, "people/123", r.ResourcePath())

}
