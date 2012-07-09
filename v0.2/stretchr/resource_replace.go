package stretchr

// Replace saves a new version of an existing resource in Stretchr.
//
// Regardless of the remote resource before, it will exactly match the state in this
// resource after this operation.  I.e. the resource will be replaced.
//
// To only partially update remote resources, see Resource.Update().
func (r *Resource) Replace() error {

	json, jsonErr := toJson(r.data)

	if jsonErr != nil {
		return jsonErr
	}

	response, _, requestErr := ActiveRequester.MakeRequest(ReplaceMethod, r.AbsoluteURL(), json, r.session.PrivateKey)

	if requestErr != nil {
		return requestErr
	}

	if response.Worked {

		// update the data of the resource
		r.data = response.Data

	} else {
		return response.GetError()
	}

	// all OK
	return nil

}
