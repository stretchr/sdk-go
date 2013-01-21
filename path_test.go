package stretchr

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestPath(t *testing.T) {

	assert.Equal(t, "people/123/books/456", Path("people", "123", "books", "456"))
	assert.Equal(t, "people/123/books", Path("people", "123", "books", ""))
	assert.Equal(t, "people/123", Path("people", "123", "", ""))

}
