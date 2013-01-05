package stretchr

import (
	"github.com/stretchrcom/stew/objects"
)

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
		panic("stretchr: Array in response data when a single object was expected.")
	case nil:
		panic("stretchr: Nothing in response data when a single object was expected.")
	}

	panic("stretchr: Something unexpected in the response data when a single object was expected.")
	return nil, nil
}
