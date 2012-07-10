package stretchr

type Many struct {
	session *Session
	path    string
}

func MakeMany(session *Session, path string) *Many {
	m := new(Many)
	m.session = session
	m.path = path
	return m
}

func (m *Many) Read() (*ResourceCollection, error) {

	var resourceCollection *ResourceCollection
	response, _, requestErr := ActiveRequester.MakeRequest(ReadMethod, m.session.Url(m.path), NoBody, m.session.PublicKey, m.session.PrivateKey)

	if requestErr != nil {
		return nil, requestErr
	}

	if response.Worked {

		// get the array of data objects
		dataObjects := response.DataCollection

		// make the resources
		resourceCollection = MakeResourceCollection()
		resourceCollection.Resources = make([]*Resource, len(dataObjects))

		for objIndex, obj := range dataObjects {
			resource := m.session.MakeResource(m.path)
			resource.data = obj.(map[string]interface{})
			resourceCollection.Resources[objIndex] = resource
		}

	} else {
		return nil, response.GetError()
	}

	return resourceCollection, nil

}
