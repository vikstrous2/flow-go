// Code generated by mockery v1.0.0. DO NOT EDIT.

package mempool

import (
	flow "github.com/onflow/flow-go/model/flow"

	mock "github.com/stretchr/testify/mock"
)

// Approvals is an autogenerated mock type for the Approvals type
type Approvals struct {
	mock.Mock
}

// Add provides a mock function with given fields: approval
func (_m *Approvals) Add(approval *flow.ResultApproval) bool {
	ret := _m.Called(approval)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*flow.ResultApproval) bool); ok {
		r0 = rf(approval)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// All provides a mock function with given fields:
func (_m *Approvals) All() []*flow.ResultApproval {
	ret := _m.Called()

	var r0 []*flow.ResultApproval
	if rf, ok := ret.Get(0).(func() []*flow.ResultApproval); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*flow.ResultApproval)
		}
	}

	return r0
}

// ByID provides a mock function with given fields: approvalID
func (_m *Approvals) ByID(approvalID flow.Identifier) (*flow.ResultApproval, bool) {
	ret := _m.Called(approvalID)

	var r0 *flow.ResultApproval
	if rf, ok := ret.Get(0).(func(flow.Identifier) *flow.ResultApproval); ok {
		r0 = rf(approvalID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.ResultApproval)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(flow.Identifier) bool); ok {
		r1 = rf(approvalID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Has provides a mock function with given fields: approvalID
func (_m *Approvals) Has(approvalID flow.Identifier) bool {
	ret := _m.Called(approvalID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(flow.Identifier) bool); ok {
		r0 = rf(approvalID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Rem provides a mock function with given fields: approvalID
func (_m *Approvals) Rem(approvalID flow.Identifier) bool {
	ret := _m.Called(approvalID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(flow.Identifier) bool); ok {
		r0 = rf(approvalID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Size provides a mock function with given fields:
func (_m *Approvals) Size() uint {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}
