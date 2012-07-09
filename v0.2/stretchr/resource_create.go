package stretchr

import (
	"errors"
	"fmt"
)

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
		r.Set("~id", response.Data["~id"])

	} else {

		if len(response.Errors) > 0 {
			return errors.New(fmt.Sprintf("%s", response.Errors[0].(map[string]interface{})["Message"]))
		} else {
			return errors.New("Unknown error")
		}

	}

	// all OK
	return nil

}
