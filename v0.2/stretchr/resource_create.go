package stretchr

// Create creates a new resource in Stretchr.  The ID of the new resource will be set.
func (r *Resource) Create() error {

	json, jsonErr := toJson(r.data)

	if jsonErr != nil {
		return jsonErr
	}

	response, _, requestErr := ActiveRequester.MakeRequest(CreateMethod, r.AbsoluteURL(), json)

	if requestErr != nil {
		return requestErr
	}

	if response.Worked {

		// get the new ID
		r.Set("IDKey", response.Data["IDKey"])

	} else {
		return response.GetError()
	}

	// all OK
	return nil

}
