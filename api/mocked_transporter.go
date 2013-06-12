package api

import (
	"github.com/stretchr/testify/mock"
)

// MockedTransporter is a Transporter object that has been mocked (using github.com/stretchr/testify/mock)
// to allow for the testing of Stretchr code, without having to make real HTTP requests.
type MockedTransporter struct {
	mock.Mock
}

// MakeRequest is a mocked version of the Transporter.MakeRequest method.
func (t *MockedTransporter) MakeRequest(request *Request) (*Response, error) {
	args := t.Called(request)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Response), nil
}
