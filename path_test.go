package stretchr

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestPath(t *testing.T) {

	assert.Equal(t, "people/123/books/456", Path("people", "123", "books", "456"))

}
