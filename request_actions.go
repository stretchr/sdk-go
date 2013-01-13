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

// extractChangeInfo checks for errors and returns the change info from a response.
func extractChangeInfo(response *api.Response) (api.ChangeInfo, error) {

	responseObject := response.BodyObject()

	//return the error if there was one
	errs := GetErrorsFromResponseObject(responseObject)
	if len(errs) > 0 {
		return nil, errs[0]
	}

	return responseObject.ChangeInfo(), nil

}

// Create creates a resource.
// If the resource exists, it will be replaced.
func (r *Request) Create(resource *Resource) (api.ChangeInfo, error) {

	response, err := r.session.underlyingSession.At(r.UnderlyingRequest.Path()).Create(resource)

	if err != nil {
		return nil, err
	}

	return extractChangeInfo(response)

}

// Update updates a resource.
// If the resource does not exist, it will be Updated.
func (r *Request) Update(resource *Resource) (api.ChangeInfo, error) {

	response, err := r.session.underlyingSession.At(r.UnderlyingRequest.Path()).Update(resource)

	if err != nil {
		return nil, err
	}

	return extractChangeInfo(response)

}

// Replace replaces a resource.
// If the resource does not exist, it will be created.
func (r *Request) Replace(resource *Resource) (api.ChangeInfo, error) {

	response, err := r.session.underlyingSession.At(r.UnderlyingRequest.Path()).Replace(resource)

	if err != nil {
		return nil, err
	}

	return extractChangeInfo(response)

}

// Delete deletes one or many resources.
func (r *Request) Delete() (api.ChangeInfo, error) {
	// TODO: https://github.com/stretchrcom/stretchr-sdk-go/issues/7

	response, err := r.session.underlyingSession.At(r.UnderlyingRequest.Path()).Delete()

	if err != nil {
		return nil, err
	}

	return extractChangeInfo(response)
}

// Save creates or updates a resource.
// If the resource doesn't exist, it will be created.
// If the resource exists, it will be updated.
func (r *Request) Save(resource *Resource) (api.ChangeInfo, error) {

	response, err := r.session.underlyingSession.At(r.UnderlyingRequest.Path()).Save(resource)

	if err != nil {
		return nil, err
	}

	return extractChangeInfo(response)

}
