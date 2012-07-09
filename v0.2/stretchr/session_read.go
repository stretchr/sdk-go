package stretchr

// Read reads a single resource from Stretchr, based on the specified path and ID values.
//
// The following code will read the person with ID "123"
//  resource, err := session.Read("people", "123")
func (s *Session) Read(path, id string) (*Resource, error) {

	resource := s.MakeResource(path).SetID(id)
	response, _, requestErr := ActiveRequester.MakeRequest(ReadMethod, resource.AbsoluteURL(), NoBody, s.PublicKey, s.PrivateKey)

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
