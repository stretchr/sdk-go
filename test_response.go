package stretchr

import (
	"bytes"
	"fmt"
	"github.com/stretchrcom/sdk-go/api"
	"github.com/stretchrcom/sdk-go/common"
	"github.com/stretchrcom/signature"
	"io/ioutil"
	"net/http"
)

var TestChangeInfo = api.ChangeInfo{"~c": 1, "~u": 2, "~d": 3, "~ids": []string{"ABC", "DEF"}}

func NewTestResponse(status float64, data interface{}, errors []map[string]interface{}, context string, changeInfo api.ChangeInfo) *api.Response {

	httpResponse := new(http.Response)

	sro := map[string]interface{}{api.ResponseObjectFieldStatusCode: status,
		api.ResponseObjectFieldData:       data,
		api.ResponseObjectFieldErrors:     errors,
		api.ResponseObjectFieldChangeInfo: changeInfo,
		api.ResponseObjectFieldContext:    context}

	session := api.NewSession("project", "publicKey", "privateKey")

	responseBytes, _ := session.Codec().Marshal(sro, nil)
	httpResponse.Body = ioutil.NopCloser(bytes.NewBuffer(responseBytes))
	httpResponse.Header = make(map[string][]string)
	httpResponse.Header[common.HeaderResponseHash] = []string{signature.HashWithKeys(responseBytes, []byte("publicKey"), []byte("privateKey"))}

	response, newResponseErr := api.NewResponse(session, httpResponse)

	if newResponseErr != nil {
		panic(fmt.Sprintf("NewTestResponse: %s", newResponseErr))
	}

	return response

}
