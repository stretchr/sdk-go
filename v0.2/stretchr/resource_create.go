package stretchr

// Create creates a new resource in Stretchr.  If you do not specify an ID for this resource,
// Stretchr will generate one for you and it will be set by this method.
func (r *Resource) Create() error {

	json, jsonErr := toJson(r.data)

	if jsonErr != nil {
		return jsonErr
	}

	response, _, requestErr := ActiveRequester.MakeRequest(CreateMethod, r.AbsoluteURL(), json, r.session.PublicKey, r.session.PrivateKey)

	if requestErr != nil {
		return requestErr
	}

	if response.Worked {

		// get the new ID
		r.Set(IDKey, response.Data[IDKey])

	} else {
		return response.GetError()
	}

	// all OK
	return nil

}
