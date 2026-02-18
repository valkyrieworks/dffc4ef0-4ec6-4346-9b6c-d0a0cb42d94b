//

package simulations

import (
	vault "github.com/valkyrieworks/vault"
	mock "github.com/stretchr/testify/mock"
)

//
type PublicKey struct {
	mock.Emulate
}

//
func (_m *PublicKey) Location() vault.Location {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 vault.Location
	if rf, ok := ret.Get(0).(func() vault.Location); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(vault.Location)
		}
	}

	return r0
}

//
func (_m *PublicKey) Octets() []byte {
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
func (_m *PublicKey) Matches(_a0 vault.PublicKey) bool {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(vault.PublicKey) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

//
func (_m *PublicKey) Kind() string {
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
func (_m *PublicKey) ValidateAutograph(msg []byte, sig []byte) bool {
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
func NewPublicKey(t interface {
	mock.TestingT
	Sanitize(func())
}) *PublicKey {
	emulate := &PublicKey{}
	mock.Emulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return emulate
}
