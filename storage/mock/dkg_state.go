// Code generated by mockery v2.13.0. DO NOT EDIT.

package mock

import (
	crypto "github.com/onflow/flow-go/crypto"
	flow "github.com/onflow/flow-go/model/flow"

	mock "github.com/stretchr/testify/mock"
)

// DKGState is an autogenerated mock type for the DKGState type
type DKGState struct {
	mock.Mock
}

// GetDKGEndState provides a mock function with given fields: epochCounter
func (_m *DKGState) GetDKGEndState(epochCounter uint64) (flow.DKGEndState, error) {
	ret := _m.Called(epochCounter)

	var r0 flow.DKGEndState
	if rf, ok := ret.Get(0).(func(uint64) flow.DKGEndState); ok {
		r0 = rf(epochCounter)
	} else {
		r0 = ret.Get(0).(flow.DKGEndState)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(epochCounter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDKGStarted provides a mock function with given fields: epochCounter
func (_m *DKGState) GetDKGStarted(epochCounter uint64) (bool, error) {
	ret := _m.Called(epochCounter)

	var r0 bool
	if rf, ok := ret.Get(0).(func(uint64) bool); ok {
		r0 = rf(epochCounter)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(epochCounter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertMyBeaconPrivateKey provides a mock function with given fields: epochCounter, key
func (_m *DKGState) InsertMyBeaconPrivateKey(epochCounter uint64, key crypto.PrivateKey) error {
	ret := _m.Called(epochCounter, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64, crypto.PrivateKey) error); ok {
		r0 = rf(epochCounter, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RetrieveMyBeaconPrivateKey provides a mock function with given fields: epochCounter
func (_m *DKGState) RetrieveMyBeaconPrivateKey(epochCounter uint64) (crypto.PrivateKey, error) {
	ret := _m.Called(epochCounter)

	var r0 crypto.PrivateKey
	if rf, ok := ret.Get(0).(func(uint64) crypto.PrivateKey); ok {
		r0 = rf(epochCounter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crypto.PrivateKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(epochCounter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetDKGEndState provides a mock function with given fields: epochCounter, endState
func (_m *DKGState) SetDKGEndState(epochCounter uint64, endState flow.DKGEndState) error {
	ret := _m.Called(epochCounter, endState)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64, flow.DKGEndState) error); ok {
		r0 = rf(epochCounter, endState)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDKGStarted provides a mock function with given fields: epochCounter
func (_m *DKGState) SetDKGStarted(epochCounter uint64) error {
	ret := _m.Called(epochCounter)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(epochCounter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewDKGStateT interface {
	mock.TestingT
	Cleanup(func())
}

// NewDKGState creates a new instance of DKGState. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDKGState(t NewDKGStateT) *DKGState {
	mock := &DKGState{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
