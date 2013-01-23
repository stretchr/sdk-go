package stretchr

import (
	"github.com/stretchrcom/sdk-go/api"
	"github.com/stretchrcom/sdk-go/common"
	"github.com/stretchrcom/stew/objects"
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

func TestResource_SetID(t *testing.T) {

	resource := MakeResourceAt("people")

	assert.Equal(t, resource, resource.SetID("tyler"))

	assert.Equal(t, resource.data[common.DataFieldID], "tyler")

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

func TestResource_ID(t *testing.T) {

	id := "ABC123"

	resource := MakeResourceAt("people")

	assert.Equal(t, NoID, resource.ID())

	resource.SetID(id)

	assert.Equal(t, id, resource.ID())

}

func TestResource_GetString(t *testing.T) {

	resource := MakeResourceAt("people")
	resource.Set("name", "Mat")

	assert.Equal(t, "Mat", resource.GetString("name"))

	resource.Set("name", 25)
	assert.Panics(t, func() {
		resource.GetString("name")
	})

}

func TestResource_GetBool(t *testing.T) {

	resource := MakeResourceAt("people")
	resource.Set("key", true)

	assert.Equal(t, true, resource.GetBool("key"))

	resource.Set("key", 25)
	assert.Panics(t, func() {
		resource.GetBool("key")
	})

}

func TestResource_GetNumber(t *testing.T) {

	resource := MakeResourceAt("people")
	resource.Set("key", float64(123))

	assert.Equal(t, 123, resource.GetNumber("key"))

	resource.Set("key", "not a number")
	assert.Panics(t, func() {
		resource.GetNumber("key")
	})

}

func TestResource_GetTime(t *testing.T) {

	resource := MakeResourceAt("people")
	atime := float64(time.Now().Unix())
	resource.Set("key", atime)

	assert.Equal(t, atime, float64(resource.GetTime("key").Unix()))

	resource.Set("key", "not a time")
	assert.Panics(t, func() {
		resource.GetTime("key")
	})

}

func TestResource_Set_NumberConversion(t *testing.T) {

	resource := MakeResourceAt("people")

	var one int = 1
	var two int8 = 2
	var three int16 = 3
	var four int32 = 4
	var five int64 = 5
	var six uint = 6
	var seven uint8 = 7
	var eight uint16 = 8
	var nine uint32 = 9
	var ten uint64 = 10
	var eleven float32 = 11.0
	var twelve float64 = 12.0

	assert.Equal(t, float64(one), resource.Set("number", one).Get("number"))
	assert.Equal(t, float64(two), resource.Set("number", two).Get("number"))
	assert.Equal(t, float64(three), resource.Set("number", three).Get("number"))
	assert.Equal(t, float64(four), resource.Set("number", four).Get("number"))
	assert.Equal(t, float64(five), resource.Set("number", five).Get("number"))
	assert.Equal(t, float64(six), resource.Set("number", six).Get("number"))
	assert.Equal(t, float64(seven), resource.Set("number", seven).Get("number"))
	assert.Equal(t, float64(eight), resource.Set("number", eight).Get("number"))
	assert.Equal(t, float64(nine), resource.Set("number", nine).Get("number"))
	assert.Equal(t, float64(ten), resource.Set("number", ten).Get("number"))
	assert.Equal(t, float64(eleven), resource.Set("number", eleven).Get("number"))
	assert.Equal(t, float64(twelve), resource.Set("number", twelve).Get("number"))

}
