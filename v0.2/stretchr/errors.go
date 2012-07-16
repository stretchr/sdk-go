package stretchr

import (
	"errors"
)

// NotFound is an error that is used when a resource could not
// be found.
var NotFound error = errors.New("Resource not found.")
