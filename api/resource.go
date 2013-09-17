package api

import (
	"github.com/stretchr/objx"
)

// Resource describes objects that can be used in conjunction with Stretchr services.
type Resource interface {

	// ResourceData gets the data for this Resource.
	ResourceData() objx.Map

	// ID gets the ID for this Resource.
	ID() string

	// SetID sets the ID for this Resource and returns the Resource for chaining.
	SetID(string) Resource
}
