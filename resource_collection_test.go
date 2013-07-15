package stretchr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewResourceCollection(t *testing.T) {

}

func TestResourceCollection_IsEmpty(t *testing.T) {

	resource := MakeResourceAt("test")
	resources := []*Resource{resource}
	rc := NewResourceCollection(resources)
	assert.Equal(t, rc.IsEmpty(), false)

}
