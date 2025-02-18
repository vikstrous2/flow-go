// Code generated by mockery v2.13.0. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"
)

// MatchingCore is an autogenerated mock type for the MatchingCore type
type MatchingCore struct {
	mock.Mock
}

// OnBlockFinalization provides a mock function with given fields:
func (_m *MatchingCore) OnBlockFinalization() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProcessReceipt provides a mock function with given fields: receipt
func (_m *MatchingCore) ProcessReceipt(receipt *flow.ExecutionReceipt) error {
	ret := _m.Called(receipt)

	var r0 error
	if rf, ok := ret.Get(0).(func(*flow.ExecutionReceipt) error); ok {
		r0 = rf(receipt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewMatchingCoreT interface {
	mock.TestingT
	Cleanup(func())
}

// NewMatchingCore creates a new instance of MatchingCore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMatchingCore(t NewMatchingCoreT) *MatchingCore {
	mock := &MatchingCore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
