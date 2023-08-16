// Code generated by mockery. DO NOT EDIT.

package users

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockFindUserProcessor is an autogenerated mock type for the FindUserProcessor type
type MockFindUserProcessor struct {
	mock.Mock
}

// ProcessRequest provides a mock function with given fields: ctx
func (_m *MockFindUserProcessor) ProcessRequest(ctx context.Context) (*[]UserResponse, error) {
	ret := _m.Called(ctx)

	var r0 *[]UserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*[]UserResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *[]UserResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]UserResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockFindUserProcessor creates a new instance of MockFindUserProcessor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFindUserProcessor(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFindUserProcessor {
	mock := &MockFindUserProcessor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
