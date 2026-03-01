//

package simulations

import (
	security "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	mock "github.com/stretchr/testify/mock"
)

//
type ClusterValidator struct {
	mock.Simulate
}

//
func (_m *ClusterValidator) Add(key security.PublicToken, artifact []byte, signing []byte) error {
	ret := _m.Called(key, artifact, signing)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(security.PublicToken, []byte, []byte) error); ok {
		r0 = rf(key, artifact, signing)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *ClusterValidator) Validate() (bool, []bool) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 bool
	var r1 []bool
	if rf, ok := ret.Get(0).(func() (bool, []bool)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func() []bool); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]bool)
		}
	}

	return r0, r1
}

//
//
func FreshClusterValidator(t interface {
	mock.TestingT
	Sanitize(func())
}) *ClusterValidator {
	simulate := &ClusterValidator{}
	mock.Simulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return simulate
}
