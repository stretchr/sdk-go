package stretchr

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestMakeResourceCollection(t *testing.T) {

}

func TestResourceCollection_IsEmpty(t *testing.T) {

	resource := MakeResourceAt("test")
	resources := []*Resource{resource}
	rc := MakeResourceCollection(resources)
	assert.Equal(t, rc.IsEmpty(), false)

}
