package stretchr

import (
	"testing"
)

func TestMakeResourceCollection(t *testing.T) {

	c := MakeResourceCollection()

	if c == nil {
		t.Error("MakeResourceCollection shouldn't return nil")
	}

}
