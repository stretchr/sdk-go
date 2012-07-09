package stretchr

// Update saves changes to an existing resource in Stretchr.  Only fields present in the resource
// will be updated, other fields (i.e. ones not mentioned in this resource but present in the remote resource)
// will be left unchanged.
//
// See also Resource.Replace().
func (r *Resource) Update() error {

	json, jsonErr := toJson(r.data)

	if jsonErr != nil {
		return jsonErr
	}

	response, _, requestErr := ActiveRequester.MakeRequest(UpdateMethod, r.AbsoluteURL(), json, r.session.PublicKey, r.session.PrivateKey)

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
