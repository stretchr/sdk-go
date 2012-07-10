package stretchr

import (
	"fmt"
	"net/url"
	"strings"
)

// Many provides the ability to work on multiple resources, such as finding collections and deleting collections etc.
type Many struct {
	session    *Session
	path       string
	parameters url.Values
}

// makeMany makes a new Many object with the given session and path.
//
// Users should create a Many object using Session.Many(path).
func makeMany(session *Session, path string) *Many {
	m := new(Many)
	m.session = session
	m.path = path
	m.parameters = make(url.Values)
	return m
}

// Path gets the URL segment and query of the request that will be made.
func (m *Many) Path() string {
	query := m.parameters.Encode()

	if len(query) > 0 {
		return fmt.Sprintf("%s?%s", m.path, query)
	}

	return m.path

}

/*
	Parameters
*/

// Parameters gets the url.Values object that holds the parameters.
func (m *Many) Parameters() url.Values {
	return m.parameters
}

// SetParameter sets a parameter on the request that will be made.
func (m *Many) SetParameter(key, value string) *Many {
	m.parameters.Set(key, value)
	return m
}

// RemoveParameter removes a parameter from the request.
func (m *Many) RemoveParameter(key string) *Many {
	delete(m.parameters, key)
	return m
}

// Limit specifies the maximum number of Resources to return.
//
// Instead of using Limit and Skip directly, its easier to use the Page function.
func (m *Many) Limit(limit int) *Many {
	return m.SetParameter(LimitKey, fmt.Sprintf("%d", limit))
}

// Skip specifies how many records to ignore before it starts collecting them.
//
// Instead of using Limit and Skip directly, its easier to use the Page function.
func (m *Many) Skip(skip int) *Many {
	if skip > 0 {
		return m.SetParameter(SkipKey, fmt.Sprintf("%d", skip))
	}
	return m.RemoveParameter(SkipKey)
}

// Page specifies the page number and size of the resources to get.
//
// For example, Page(1, 10) will get the first page of 10 records.
func (m *Many) Page(pageNumber, pageSize int) *Many {
	return m.Limit(pageSize).Skip(pageSize * (pageNumber - 1))
}

// Order specifies the order in which resources should be returned.
//
// To order first by age (oldest first), then by name:
//  .Order("-age", "name")
func (m *Many) Order(keys ...string) *Many {
	return m.SetParameter(OrderKey, strings.Join(keys, ","))
}

/*
	Data operations
*/

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
