package api

import (
	"github.com/stretchrcom/stew/objects"
)

// Resource describes objects that can be used in conjunction with Stretchr services.
type Resource interface {

	// ResourceData gets the data for this Resource.
	ResourceData() objects.Map
}
