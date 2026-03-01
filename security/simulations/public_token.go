//

package simulations

import (
	security "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	mock "github.com/stretchr/testify/mock"
)

//
type PublicToken struct {
	mock.Simulate
}

//
func (_m *PublicToken) Location() security.Location {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 security.Location
	if rf, ok := ret.Get(0).(func() security.Location); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(security.Location)
		}
	}

	return r0
}

//
func (_m *PublicToken) Octets() []byte {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

//
func (_m *PublicToken) Matches(_a0 security.PublicToken) bool {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(security.PublicToken) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

//
func (_m *PublicToken) Kind() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

//
func (_m *PublicToken) ValidateSigning(msg []byte, sig []byte) bool {
	ret := _m.Called(msg, sig)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func([]byte, []byte) bool); ok {
		r0 = rf(msg, sig)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

//
//
func FreshPublicToken(t interface {
	mock.TestingT
	Sanitize(func())
}) *PublicToken {
	simulate := &PublicToken{}
	mock.Simulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return simulate
}
