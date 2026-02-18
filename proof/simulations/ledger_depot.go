//

package simulations

import (
	kinds "github.com/valkyrieworks/kinds"
	mock "github.com/stretchr/testify/mock"
)

//
type LedgerDepot struct {
	mock.Emulate
}

//
func (_m *LedgerDepot) Level() int64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

//
func (_m *LedgerDepot) ImportLedgerEndorse(level int64) *kinds.Endorse {
	ret := _m.Called(level)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.Endorse
	if rf, ok := ret.Get(0).(func(int64) *kinds.Endorse); ok {
		r0 = rf(level)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.Endorse)
		}
	}

	return r0
}

//
func (_m *LedgerDepot) ImportLedgerMeta(level int64) *kinds.LedgerMeta {
	ret := _m.Called(level)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.LedgerMeta
	if rf, ok := ret.Get(0).(func(int64) *kinds.LedgerMeta); ok {
		r0 = rf(level)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.LedgerMeta)
		}
	}

	return r0
}

//
//
func NewLedgerDepot(t interface {
	mock.TestingT
	Sanitize(func())
}) *LedgerDepot {
	emulate := &LedgerDepot{}
	mock.Emulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return emulate
}
