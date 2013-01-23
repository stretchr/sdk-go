package stretchr

import (
	"github.com/stretchrcom/sdk-go/api"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestTestResponse_NewTestResponse(t *testing.T) {

	status := float64(201)
	data := map[string]interface{}{"name": "Mat"}
	errors := []map[string]interface{}{map[string]interface{}{api.ResponseObjectFieldErrorsMessage: "An Error"}}
	context := "context"
	changeInfo := TestChangeInfo

	r := NewTestResponse(status, data, errors, context, changeInfo)

	assert.Equal(t, 201, r.BodyObject().StatusCode())
	assert.Equal(t, context, r.BodyObject().Context())
	assert.Equal(t, data["name"], r.BodyObject().Data().(map[string]interface{})["name"])
	if assert.Equal(t, 1, len(r.BodyObject().Errors()), "There should be one error.") {
		assert.Equal(t, "An Error", r.BodyObject().Errors()[0])
	}
	if assert.NotNil(t, r.BodyObject().ChangeInfo()) {

		assert.Equal(t, 1, r.BodyObject().ChangeInfo().Created())
		assert.Equal(t, 2, r.BodyObject().ChangeInfo().Updated())
		assert.Equal(t, 3, r.BodyObject().ChangeInfo().Deleted())

		ids := r.BodyObject().ChangeInfo().IDs()

		if assert.Equal(t, 2, len(ids)) {
			assert.Equal(t, "ABC", ids[0])
			assert.Equal(t, "DEF", ids[1])
		}

	}

}
