package stretchr

import (
	"github.com/stretchrcom/testify/assert"
	"net/http"
	"testing"
)

func TestResponse_NewResponse(t *testing.T) {

	httpResponse := new(http.Response)
	response := NewResponse(httpResponse)

	assert.Equal(t, response.HttpResponse, httpResponse)

}
