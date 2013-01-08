package stretchr

import (
	"github.com/stretchrcom/stew/objects"
)

// LoadOne loads a resource from Stretchr with the given path.
func (s *Session) LoadOne(path string) (*Resource, error) {

	response, err := s.session.At(path).Read()

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
		resource := MakeResourceAt(path)
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
func (s *Session) LoadMany(path string) ([]*Resource, error) {

	response, err := s.session.At(path).Read()

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

		return nil, nil
	case map[string]interface{}:
		return nil, ErrArrayObjectExpectedButGotSingleObject
	case nil:
		return nil, ErrArrayObjectExpectedButGotNil
	}

	return nil, ErrArrayObjectExpectedButGotSomethingElse

}
