package stretchr

import (
	"github.com/stretchrcom/sdk-go/api"
	"github.com/stretchrcom/sdk-go/common"
	"github.com/stretchrcom/stew/objects"
)

// ReadOne loads a resource from Stretchr with the given path.
func (r *Request) ReadOne() (*Resource, error) {

	response, err := r.UnderlyingRequest.Read()

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

// ReadMany loads many resources from Stretchr with the given path.
func (r *Request) ReadMany() (*ResourceCollection, error) {

	response, err := r.UnderlyingRequest.Read()

	if err != nil {
		return nil, err
	}

	responseObject := response.BodyObject()

	// return the error if there was one
	errs := GetErrorsFromResponseObject(responseObject)
	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resourceArray, exists := responseObject.Data().(map[string]interface{})["~i"].([]interface{}); exists {
		resources := make([]*Resource, len(resourceArray))

		// populate the resources
		for resIndex, responseData := range resourceArray {
			resource := MakeResourceAt(r.UnderlyingRequest.Path())
			resource.data = objects.Map(responseData.(map[string]interface{})).Copy()
			resources[resIndex] = resource
		}

		resourceCollection := MakeResourceCollection(resources)

		return resourceCollection, nil
	} else {
		return nil, ErrArrayObjectExpectedButGotSomethingElse
	}

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
func (r *Request) Create(resource api.Resource) (api.ChangeInfo, error) {

	response, err := r.UnderlyingRequest.Create(resource)

	if err != nil {
		return nil, err
	}

	changeInfo, err := extractChangeInfo(response)

	if err != nil {
		return nil, err
	}

	if changeInfo.Created() == 1 {
		resource.SetID(changeInfo.Deltas()[0][common.DataFieldID].(string))
	}

	resource.ResourceData().MergeHere(changeInfo.Deltas()[0])

	return changeInfo, nil

}

// CreateMany creates many resources.
// If a resource exists, it will be replaced.
func (r *Request) CreateMany(resourceCollection *ResourceCollection) (api.ChangeInfo, error) {

	// We have to manuall repackage the collection data so the api.CreateMany will accept it
	var data []api.Resource

	for _, resource := range resourceCollection.Resources {
		data = append(data, resource)
	}

	response, err := r.UnderlyingRequest.CreateMany(data)

	if err != nil {
		return nil, err
	}

	changeInfo, err := extractChangeInfo(response)

	if err != nil {
		return nil, err
	}

	for index, resource := range resourceCollection.Resources {
		resource.ResourceData().MergeHere(changeInfo.Deltas()[index])
	}

	return changeInfo, nil

}

// Update updates a resource.
// If the resource does not exist, it will be Updated.
func (r *Request) Update(resource api.Resource) (api.ChangeInfo, error) {

	response, err := r.UnderlyingRequest.Update(resource)

	if err != nil {
		return nil, err
	}

	changeInfo, err := extractChangeInfo(response)

	if err != nil {
		return nil, err
	}

	resource.ResourceData().MergeHere(changeInfo.Deltas()[0])

	return changeInfo, nil

}

// Replace replaces a resource.
// If the resource does not exist, it will be created.
func (r *Request) Replace(resource api.Resource) (api.ChangeInfo, error) {

	response, err := r.UnderlyingRequest.Replace(resource)

	if err != nil {
		return nil, err
	}

	changeInfo, err := extractChangeInfo(response)

	if err != nil {
		return nil, err
	}

	resource.ResourceData().MergeHere(changeInfo.Deltas()[0])

	return changeInfo, nil

}

// Delete deletes one or many resources.
func (r *Request) Delete() (api.ChangeInfo, error) {
	// TODO: https://github.com/stretchrcom/sdk-go/issues/7

	response, err := r.UnderlyingRequest.Delete()

	if err != nil {
		return nil, err
	}

	return extractChangeInfo(response)
}

// Save creates or updates a resource.
// If the resource doesn't exist, it will be created.
// If the resource exists, it will be updated.
func (r *Request) Save(resource api.Resource) (api.ChangeInfo, error) {

	response, err := r.UnderlyingRequest.Save(resource)

	if err != nil {
		return nil, err
	}

	changeInfo, err := extractChangeInfo(response)

	if err != nil {
		return nil, err
	}

	resource.ResourceData().MergeHere(changeInfo.Deltas()[0])

	return changeInfo, nil

}
