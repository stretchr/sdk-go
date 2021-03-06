package stretchr

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrors_GetErrorsFromResponseObject(t *testing.T) {

	response := NewTestResponse(500, nil, []map[string]interface{}{map[string]interface{}{"~message": "Oops One"}, map[string]interface{}{"~message": "Oops Two"}}, "", nil)

	errs := GetErrorsFromResponseObject(response.BodyObject())

	if assert.Equal(t, 2, len(errs)) {
		assert.Equal(t, "Oops One", fmt.Sprintf("%s", errs[0]))
		assert.Equal(t, "Oops Two", fmt.Sprintf("%s", errs[1]))
	}

}

func TestErrors_GetErrorFromResponseObject(t *testing.T) {

	response := NewTestResponse(500, nil, []map[string]interface{}{map[string]interface{}{"~message": "Oops One"}, map[string]interface{}{"~message": "Oops Two"}}, "", nil)

	assert.Equal(t, "Oops One", fmt.Sprintf("%s", GetErrorFromResponseObject(response.BodyObject())))

}
