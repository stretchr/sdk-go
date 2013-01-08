package stretchr

import (
	"github.com/stretchrcom/stew/objects"
	"github.com/stretchrcom/stretchr-sdk-go/api"
)

// LoadOne loads a resource from Stretchr with the given path.
func (r *Request) LoadOne() (*Resource, error) {

	response, err := r.session.underlyingSession.At(r.UnderlyingRequest.Path()).Read()

	if err != nil {
		return nil, err
	}

	responseObject := response.BodyObject()

	// return the error if there was one
	errs := GetErrorsFromResponseObject(responseObject)
	if len(errs) > 0 {
		return nil, errs[0]
	}

	switch responseObject.Data().(type) {
	case map[string]interface{}:
		resource := MakeResourceAt(r.UnderlyingRequest.Path())
		resource.data = objects.Map(responseObject.Data().(map[string]interface{})).Copy()
		return resource, nil
	case []interface{}:
		return nil, ErrSingleObjectExpectedButGotArray
	case nil:
		return nil, ErrSingleObjectExpectedButGotNil
	}

	return nil, ErrSingleObjectExpectedButGotSomethingElse
}

// LoadMany loads many resources from Stretchr with the given path.
func (r *Request) LoadMany() (*ResourceCollection, error) {

	response, err := r.session.underlyingSession.At(r.UnderlyingRequest.Path()).Read()

	if err != nil {
		return nil, err
	}

	responseObject := response.BodyObject()

	// return the error if there was one
	errs := GetErrorsFromResponseObject(responseObject)
	if len(errs) > 0 {
		return nil, errs[0]
	}

	switch responseObject.Data().(type) {
	case []interface{}:

		data := responseObject.Data().([]interface{})
		resources := make([]*Resource, len(data))

		// populate the resources
		for resIndex, responseData := range data {
			resource := MakeResourceAt(r.UnderlyingRequest.Path())
			resource.data = objects.Map(responseData.(map[string]interface{})).Copy()
			resources[resIndex] = resource
		}

		resourceCollection := MakeResourceCollection(resources)

		return resourceCollection, nil
	case map[string]interface{}:
		return nil, ErrArrayObjectExpectedButGotSingleObject
	case nil:
		return nil, ErrArrayObjectExpectedButGotNil
	}

	return nil, ErrArrayObjectExpectedButGotSomethingElse

}

func (r *Request) Delete() (api.ChangeInfo, error) {

	response, err := r.session.underlyingSession.At(r.UnderlyingRequest.Path()).Delete()

	if err != nil {
		return nil, err
	}

	responseObject := response.BodyObject()

	// return the error if there was one
	errs := GetErrorsFromResponseObject(responseObject)
	if len(errs) > 0 {
		return nil, errs[0]
	}

	return responseObject.ChangeInfo(), nil
}
