package stretchr

// Delete deletes a single resource specified by the path and id arguments.
//
// For example, the following code will delete a people record with ID 'simon':
//  deleteErr := session.Delete("people", "simon")
func (s *Session) Delete(path, id string) error {

	resource := s.MakeResource(path).SetID(id)
	response, _, requestErr := ActiveRequester.MakeRequest(DeleteMethod, resource.AbsoluteURL(), NoBody, s.PublicKey, s.PrivateKey)

	if requestErr != nil {
		return requestErr
	}

	if !response.Worked {
		return response.GetError()
	}

	// success
	return nil

}
