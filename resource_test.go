package stretchr

import (
	"github.com/stretchrcom/stew/objects"
	"github.com/stretchrcom/stretchr-sdk-go/api"
	"github.com/stretchrcom/stretchr-sdk-go/common"
	"github.com/stretchrcom/testify/assert"
	"testing"
	"time"
)

func TestResource_Interface(t *testing.T) {

	assert.Implements(t, (*api.Resource)(nil), new(Resource))

}

func TestMakeResource(t *testing.T) {

	r := MakeResourceAt("people/123")
	r.data["name"] = "Mat"

	assert.Equal(t, r.data[common.DataFieldID], "123", "MakeResourceAt should set the ID in the data")

	if assert.NotNil(t, r) {
		assert.Equal(t, "people/123", r.ResourcePath())
		assert.Equal(t, "Mat", r.ResourceData()["name"])
	}

}

func TestResource_ResourcePath_GetsIDFromBody(t *testing.T) {

	r := MakeResourceAt("people")
	r.data[common.DataFieldID] = "123"

	assert.Equal(t, "people/123", r.ResourcePath())

}
func TestResource_ResourcePath_ReplacesIDFromBody(t *testing.T) {

	r := MakeResourceAt("people/abc")
	r.data[common.DataFieldID] = "123"

	assert.Equal(t, "people/123", r.ResourcePath())

}

func TestResource_Set(t *testing.T) {

	resource := MakeResourceAt("people")

	atime := time.Now()

	assert.Equal(t, resource.Set("name", "Tyler"), resource)

	resource.Set("name", "Tyler").
		Set("state", "ut").
		Set("abool", true).
		Set("now", atime).
		Set("number", 123)

	assert.Equal(t, resource.data["name"], "Tyler")
	assert.Equal(t, resource.data["state"], "ut")
	assert.Equal(t, resource.data["abool"], true)
	assert.Equal(t, resource.data["now"], atime)
	assert.Equal(t, resource.data["number"], 123)

}

func TestResource_Get(t *testing.T) {

	resource := MakeResourceAt("people")

	atime := time.Now()

	assert.Equal(t, resource.Set("name", "Tyler"), resource)

	resource.Set("name", "Tyler").
		Set("state", "ut").
		Set("abool", true).
		Set("now", atime).
		Set("number", 123)

	assert.Equal(t, resource.Get("name"), "Tyler")
	assert.Equal(t, resource.Get("state"), "ut")
	assert.Equal(t, resource.Get("abool"), true)
	assert.Equal(t, resource.Get("now"), atime)
	assert.Equal(t, resource.Get("number"), 123)

	deep := objects.Map{"name": objects.Map{"first": "Mat", "last": "Ryer"}}

	resource.data["deep"] = deep
	assert.Equal(t, "Mat", resource.Get("deep.name.first"))
	assert.Equal(t, "Ryer", resource.Get("deep.name.last"))

}

func TestResource_Save(t *testing.T) {

	//resource := MakeResourceAt("people")
	//resource.

}
