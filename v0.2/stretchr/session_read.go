package stretchr

func (s *Session) Read(path, id string) (*Resource, error) {

	resource := s.MakeResource(path).SetID(id)
	response, _, requestErr := ActiveRequester.MakeRequest(ReadMethod, resource.AbsoluteURL(), NoBody)

	if requestErr != nil {
		return resource, requestErr
	}

	if response.Worked {

		// update the data of the resource
		resource.data = response.Data

	} else {
		return nil, response.GetError()
	}

	return resource, nil
}
