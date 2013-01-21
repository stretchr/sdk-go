package stretchr

/*

	This file demonstrates how you might write strongly-typed resources
	to make your Stretchr code easier (and compiler friendly).

*/

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

// PersonResource represents an example of a strongly-typed Resource object.
type PersonResource struct {
	Resource
}

func MakePersonResource(id string) *PersonResource {
	p := new(PersonResource)
	p.Resource = *MakeResourceAt(Path("people", id))
	return p
}

// Name is an example of a strongly typed and managed getter.
func (p *PersonResource) Name() string {
	return p.GetString("name")
}

// SetName is an example of a srongly typed and managed setter.
func (p *PersonResource) SetName(value string) *PersonResource {
	p.Set("name", value)
	return p
}

// Age is an example of a strongly typed and managed getter that does some
// casting between int, and the underlying float64 number type.
func (p *PersonResource) Age() int {
	return int(p.GetNumber("age"))
}

// SetAge is an exmaple of a strongly typed and managed setter that does some
// casting between int, and the underlying float64 number type.
func (p *PersonResource) SetAge(value int) *PersonResource {
	p.SetNumber("age", float64(value))
	return p
}

/*

	This is how you create an instance of your strongly
	typed resource.

*/
func TestStrongResources_MakingPersonResource(t *testing.T) {

	p := new(PersonResource)
	p.Resource = *MakeResourceAt(Path("people", "123"))

	assert.Equal(t, "people/123", p.ResourcePath())

}

/*

	This is how you would use your strongly-typed and managed getters
	and setters.

*/
func TestStrongResources_StrongGettersAndSetters(t *testing.T) {

	p := new(PersonResource)
	p.Resource = *MakeResourceAt("people/123")

	p.SetName("Mat").SetAge(30)
	assert.Equal(t, "Mat", p.Name())
	assert.Equal(t, int(30), p.Age())

	// ensure also that the actual data is set in the resource
	assert.Equal(t, "Mat", p.Resource.ResourceData().Get("name"))
	assert.Equal(t, 30, p.Resource.ResourceData().Get("age"))

}
