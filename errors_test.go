package stretchr

import (
	"fmt"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestErrors_GetErrorsFromResponseObject(t *testing.T) {

	response := NewTestResponse(500, nil, []map[string]interface{}{map[string]interface{}{"m": "Oops One"}, map[string]interface{}{"m": "Oops Two"}}, "", nil)

	errs := GetErrorsFromResponseObject(response.BodyObject())

	if assert.Equal(t, 2, len(errs)) {
		assert.Equal(t, "Oops One", fmt.Sprintf("%s", errs[0]))
		assert.Equal(t, "Oops Two", fmt.Sprintf("%s", errs[1]))
	}

}
