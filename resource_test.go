package stretchr

import (
	"github.com/stretchrcom/stretchr-sdk-go/api"
	"github.com/stretchrcom/stretchr-sdk-go/common"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestResource_Interface(t *testing.T) {

	assert.Implements(t, (*api.Resource)(nil), new(Resource))

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

func TestResource_ResourcePath_GetsIDFromBody(t *testing.T) {

	r := MakeResourceAt("people")
	r.Data[common.DataFieldID] = "123"

	assert.Equal(t, "people/123", r.ResourcePath())

}
func TestResource_ResourcePath_ReplacesIDFromBody(t *testing.T) {

	r := MakeResourceAt("people/abc")
	r.Data[common.DataFieldID] = "123"

	assert.Equal(t, "people/123", r.ResourcePath())

}
