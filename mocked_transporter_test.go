package stretchr

import (
	"github.com/stretchrcom/testify/mock"
)

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
