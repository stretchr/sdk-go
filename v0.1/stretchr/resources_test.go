package stretchr

import (
	"testing"
)

func TestMakeResource(t *testing.T) {

	var r Resource = MakeResource()

	if r == nil {
		t.Error("MakeResource() should return new Resource.")
	}

}
