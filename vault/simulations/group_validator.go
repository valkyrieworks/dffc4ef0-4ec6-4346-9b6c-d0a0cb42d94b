//

package simulations

import (
	vault "github.com/valkyrieworks/vault"
	mock "github.com/stretchr/testify/mock"
)

//
type GroupValidator struct {
	mock.Emulate
}

//
func (_m *GroupValidator) Add(key vault.PublicKey, signal []byte, autograph []byte) error {
	ret := _m.Called(key, signal, autograph)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(vault.PublicKey, []byte, []byte) error); ok {
		r0 = rf(key, signal, autograph)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *GroupValidator) Validate() (bool, []bool) {
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
func NewGroupValidator(t interface {
	mock.TestingT
	Sanitize(func())
}) *GroupValidator {
	emulate := &GroupValidator{}
	mock.Emulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return emulate
}
