package stretchr

// Many provides the ability to work on multiple resources, such as finding collections and deleting collections etc.
type Many struct {
	session *Session
	path    string
}

// MakeMany makes a new Many object with the given session and path.
//
// Best practice is to use Session.Many(path) instead.
func MakeMany(session *Session, path string) *Many {
	m := new(Many)
	m.session = session
	m.path = path
	return m
}

// Read reads many resources from Stretchr based on the configuration in this Many object.
//
// The following code will read the first 10 people:
//  people, err := session.Many("people").Limit(10).Read()
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
