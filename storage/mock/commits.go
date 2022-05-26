// Code generated by mockery v2.12.1. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"

	storage "github.com/onflow/flow-go/storage"

	testing "testing"
)

// Commits is an autogenerated mock type for the Commits type
type Commits struct {
	mock.Mock
}

// BatchStore provides a mock function with given fields: blockID, commit, batch
func (_m *Commits) BatchStore(blockID flow.Identifier, commit flow.StateCommitment, batch storage.BatchStorage) error {
	ret := _m.Called(blockID, commit, batch)

	var r0 error
	if rf, ok := ret.Get(0).(func(flow.Identifier, flow.StateCommitment, storage.BatchStorage) error); ok {
		r0 = rf(blockID, commit, batch)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ByBlockID provides a mock function with given fields: blockID
func (_m *Commits) ByBlockID(blockID flow.Identifier) (flow.StateCommitment, error) {
	ret := _m.Called(blockID)

	var r0 flow.StateCommitment
	if rf, ok := ret.Get(0).(func(flow.Identifier) flow.StateCommitment); ok {
		r0 = rf(blockID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.StateCommitment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(flow.Identifier) error); ok {
		r1 = rf(blockID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: blockID, commit
func (_m *Commits) Store(blockID flow.Identifier, commit flow.StateCommitment) error {
	ret := _m.Called(blockID, commit)

	var r0 error
	if rf, ok := ret.Get(0).(func(flow.Identifier, flow.StateCommitment) error); ok {
		r0 = rf(blockID, commit)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCommits creates a new instance of Commits. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommits(t testing.TB) *Commits {
	mock := &Commits{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
