//

package simulations

import (
	status "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	mock "github.com/stretchr/testify/mock"

	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
type ProofHub struct {
	mock.Simulate
}

//
func (_m *ProofHub) AppendProof(_a0 kinds.Proof) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(kinds.Proof) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *ProofHub) InspectProof(_a0 kinds.ProofCatalog) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(kinds.ProofCatalog) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *ProofHub) AwaitingProof(maximumOctets int64) ([]kinds.Proof, int64) {
	ret := _m.Called(maximumOctets)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 []kinds.Proof
	var r1 int64
	if rf, ok := ret.Get(0).(func(int64) ([]kinds.Proof, int64)); ok {
		return rf(maximumOctets)
	}
	if rf, ok := ret.Get(0).(func(int64) []kinds.Proof); ok {
		r0 = rf(maximumOctets)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]kinds.Proof)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) int64); ok {
		r1 = rf(maximumOctets)
	} else {
		r1 = ret.Get(1).(int64)
	}

	return r0, r1
}

//
func (_m *ProofHub) Revise(_a0 status.Status, _a1 kinds.ProofCatalog) {
	_m.Called(_a0, _a1)
}

//
//
func FreshProofHub(t interface {
	mock.TestingT
	Sanitize(func())
}) *ProofHub {
	simulate := &ProofHub{}
	mock.Simulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return simulate
}
