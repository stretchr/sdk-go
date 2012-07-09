package stretchr

func (r *Resource) Update() error {

	json, jsonErr := toJson(r.data)

	if jsonErr != nil {
		return jsonErr
	}

	response, _, requestErr := ActiveRequester.MakeRequest(UpdateMethod, r.AbsoluteURL(), json)

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
