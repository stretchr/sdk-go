package stretchr

/*

	This file demonstrates how you might write strongly-typed resources
	to make your Stretchr code easier (and compiler friendly).

*/

import (
	"github.com/stretchr/sdk-go/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// PersonResource represents an example of a strongly-typed Resource object.
//
// In your own code, you can follow this pattern:
//
//     type MyStrongType struct {
//       *stretchr.Resource
//     }
type PersonResource struct {
	/*

		Embedding a Resource type pointer as an anonymous field allows
		this type to act like a Resource in every necessary way.

	*/
	*Resource
}

// MakePersonResource makes a new person resource with the given ID.
//
// It also configures the actual Resource part of this object by assigning a
// new Resource to the .Resource variable.
func MakePersonResource(id string) *PersonResource {
	p := new(PersonResource)
	p.Resource = MakeResourceAt(Path("people", id))
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
	p.Set("age", value)
	return p
}

/*

	This is how you create an instance of your strongly
	typed resource.

*/
func TestStrongResources_MakingPersonResource(t *testing.T) {

	p := new(PersonResource)
	p.Resource = MakeResourceAt(Path("people", "123"))

	assert.Equal(t, "people/123", p.ResourcePath())

}

/*

	This is how you would use your strongly-typed and managed getters
	and setters.

*/
func TestStrongResources_StrongGettersAndSetters(t *testing.T) {

	p := new(PersonResource)
	p.Resource = MakeResourceAt("people/123")

	p.SetName("Mat").SetAge(30)
	assert.Equal(t, "Mat", p.Name())
	assert.Equal(t, int(30), p.Age())

	// ensure also that the actual data is set in the resource
	assert.Equal(t, "Mat", p.Resource.ResourceData().Get("name"))
	assert.Equal(t, 30, p.Resource.ResourceData().Get("age"))

}

/*

	This shows how you can use your strongly typed resource in place
	of the normal Resource object with zero effort

	The mockedTransporter stuff is just so our call to "Create" doesn't
	actually try to hit any servers.

*/
func TestStrongResources_UsingTheResource(t *testing.T) {

	// make a session object
	session := NewSession("project", "publicKey", "privateKey")

	// don't make real requests
	mockedTransporter := new(api.MockedTransporter)
	session.SetTransporter(mockedTransporter)
	mockedTransporter.On("MakeRequest", mock.Anything).Return(nil, assert.AnError)

	// make a person resource
	p := MakePersonResource("123")

	//... and use it as normal
	session.At(p.ResourcePath()).Create(p)

}

/*

	This is how, after reading a Resource, you are able to easily cast it
	to your strongly typed resource.

	We recommend you wrap this kind of thing in a function, in this case:

	    func ReadPerson(id string) *PersonResource { ... }

*/
func TestStrongResources_ReadingAResource(t *testing.T) {

	loadedResource := MakeResourceAt("people/123")
	loadedResource.Set("name", "Mat")
	loadedResource.Set("age", float64(30))

	stronglyTypedResource := MakePersonResource(loadedResource.ID())
	stronglyTypedResource.Resource = loadedResource

	assert.Equal(t, "Mat", stronglyTypedResource.Name())
	assert.Equal(t, 30, stronglyTypedResource.Age())

}
