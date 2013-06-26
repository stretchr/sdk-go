package stretchr

import (
	"bytes"
	"fmt"
	"github.com/stretchr/sdk-go/api"
	"github.com/stretchr/sdk-go/common"
	"github.com/stretchr/signature"
	"io/ioutil"
	"net/http"
)

var TestChangeInfo = api.ChangeInfo{common.ChangeInfoPublicFieldCreated: 1, common.ChangeInfoPublicFieldUpdated: 2, common.ChangeInfoPublicFieldDeleted: 3, common.ChangeInfoPublicFieldDeltas: []interface{}{map[string]interface{}{common.DataFieldID: "ABC"}, map[string]interface{}{common.DataFieldID: "DEF"}}}

func NewTestResponse(status float64, data interface{}, errors []map[string]interface{}, context string, changeInfo api.ChangeInfo) *api.Response {

	httpResponse := new(http.Response)

	sro := map[string]interface{}{common.ResponseObjectFieldStatusCode: status,
		common.ResponseObjectFieldData:       data,
		common.ResponseObjectFieldErrors:     errors,
		common.ResponseObjectFieldChangeInfo: changeInfo,
		common.ResponseObjectFieldContext:    context}

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
